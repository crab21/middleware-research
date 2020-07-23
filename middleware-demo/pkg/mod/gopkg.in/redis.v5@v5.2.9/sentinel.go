package redis

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"gopkg.in/redis.v5/internal"
	"gopkg.in/redis.v5/internal/pool"
)

//------------------------------------------------------------------------------

// FailoverOptions are used to configure a failover client and should
// be passed to NewFailoverClient.
type FailoverOptions struct {
	// The main name.
	MainName string
	// A seed list of host:port addresses of sentinel nodes.
	SentinelAddrs []string

	// Following options are copied from Options struct.

	Password string
	DB       int

	MaxRetries int

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

func (opt *FailoverOptions) options() *Options {
	return &Options{
		Addr: "FailoverClient",

		DB:       opt.DB,
		Password: opt.Password,

		MaxRetries: opt.MaxRetries,

		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,

		PoolSize:           opt.PoolSize,
		PoolTimeout:        opt.PoolTimeout,
		IdleTimeout:        opt.IdleTimeout,
		IdleCheckFrequency: opt.IdleCheckFrequency,
	}
}

// NewFailoverClient returns a Redis client that uses Redis Sentinel
// for automatic failover. It's safe for concurrent use by multiple
// goroutines.
func NewFailoverClient(failoverOpt *FailoverOptions) *Client {
	opt := failoverOpt.options()
	opt.init()

	failover := &sentinelFailover{
		mainName:    failoverOpt.MainName,
		sentinelAddrs: failoverOpt.SentinelAddrs,

		opt: opt,
	}

	client := Client{
		baseClient: baseClient{
			opt:      opt,
			connPool: failover.Pool(),

			onClose: func() error {
				return failover.Close()
			},
		},
	}
	client.cmdable.process = client.Process

	return &client
}

//------------------------------------------------------------------------------

type sentinelClient struct {
	cmdable
	baseClient
}

func newSentinel(opt *Options) *sentinelClient {
	opt.init()
	client := sentinelClient{
		baseClient: baseClient{
			opt:      opt,
			connPool: newConnPool(opt),
		},
	}
	client.cmdable = cmdable{client.Process}
	return &client
}

func (c *sentinelClient) PubSub() *PubSub {
	return &PubSub{
		base: baseClient{
			opt:      c.opt,
			connPool: pool.NewStickyConnPool(c.connPool.(*pool.ConnPool), false),
		},
	}
}

func (c *sentinelClient) GetMainAddrByName(name string) *StringSliceCmd {
	cmd := NewStringSliceCmd("SENTINEL", "get-main-addr-by-name", name)
	c.Process(cmd)
	return cmd
}

func (c *sentinelClient) Sentinels(name string) *SliceCmd {
	cmd := NewSliceCmd("SENTINEL", "sentinels", name)
	c.Process(cmd)
	return cmd
}

type sentinelFailover struct {
	mainName    string
	sentinelAddrs []string

	opt *Options

	pool     *pool.ConnPool
	poolOnce sync.Once

	mu       sync.RWMutex
	sentinel *sentinelClient
}

func (d *sentinelFailover) Close() error {
	return d.resetSentinel()
}

func (d *sentinelFailover) dial() (net.Conn, error) {
	addr, err := d.MainAddr()
	if err != nil {
		return nil, err
	}
	return net.DialTimeout("tcp", addr, d.opt.DialTimeout)
}

func (d *sentinelFailover) Pool() *pool.ConnPool {
	d.poolOnce.Do(func() {
		d.opt.Dialer = d.dial
		d.pool = newConnPool(d.opt)
	})
	return d.pool
}

func (d *sentinelFailover) MainAddr() (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Try last working sentinel.
	if d.sentinel != nil {
		addr, err := d.sentinel.GetMainAddrByName(d.mainName).Result()
		if err != nil {
			internal.Logf("sentinel: GetMainAddrByName %q failed: %s", d.mainName, err)
			d._resetSentinel()
		} else {
			addr := net.JoinHostPort(addr[0], addr[1])
			internal.Logf("sentinel: %q addr is %s", d.mainName, addr)
			return addr, nil
		}
	}

	for i, sentinelAddr := range d.sentinelAddrs {
		sentinel := newSentinel(&Options{
			Addr: sentinelAddr,

			DialTimeout:  d.opt.DialTimeout,
			ReadTimeout:  d.opt.ReadTimeout,
			WriteTimeout: d.opt.WriteTimeout,

			PoolSize:    d.opt.PoolSize,
			PoolTimeout: d.opt.PoolTimeout,
			IdleTimeout: d.opt.IdleTimeout,
		})
		mainAddr, err := sentinel.GetMainAddrByName(d.mainName).Result()
		if err != nil {
			internal.Logf("sentinel: GetMainAddrByName %q failed: %s", d.mainName, err)
			sentinel.Close()
			continue
		}

		// Push working sentinel to the top.
		d.sentinelAddrs[0], d.sentinelAddrs[i] = d.sentinelAddrs[i], d.sentinelAddrs[0]

		d.setSentinel(sentinel)
		addr := net.JoinHostPort(mainAddr[0], mainAddr[1])
		internal.Logf("sentinel: %q addr is %s", d.mainName, addr)
		return addr, nil
	}

	return "", errors.New("redis: all sentinels are unreachable")
}

func (d *sentinelFailover) setSentinel(sentinel *sentinelClient) {
	d.discoverSentinels(sentinel)
	d.sentinel = sentinel
	go d.listen(sentinel)
}

func (d *sentinelFailover) resetSentinel() error {
	d.mu.Lock()
	err := d._resetSentinel()
	d.mu.Unlock()
	return err
}

func (d *sentinelFailover) _resetSentinel() error {
	var err error
	if d.sentinel != nil {
		err = d.sentinel.Close()
		d.sentinel = nil
	}
	return err
}

func (d *sentinelFailover) discoverSentinels(sentinel *sentinelClient) {
	sentinels, err := sentinel.Sentinels(d.mainName).Result()
	if err != nil {
		internal.Logf("sentinel: Sentinels %q failed: %s", d.mainName, err)
		return
	}
	for _, sentinel := range sentinels {
		vals := sentinel.([]interface{})
		for i := 0; i < len(vals); i += 2 {
			key := vals[i].(string)
			if key == "name" {
				sentinelAddr := vals[i+1].(string)
				if !contains(d.sentinelAddrs, sentinelAddr) {
					internal.Logf(
						"sentinel: discovered new %q sentinel: %s",
						d.mainName, sentinelAddr,
					)
					d.sentinelAddrs = append(d.sentinelAddrs, sentinelAddr)
				}
			}
		}
	}
}

// closeOldConns closes connections to the old main after failover switch.
func (d *sentinelFailover) closeOldConns(newMain string) {
	// Good connections that should be put back to the pool. They
	// can't be put immediately, because pool.PopFree will return them
	// again on next iteration.
	cnsToPut := make([]*pool.Conn, 0)

	for {
		cn := d.pool.PopFree()
		if cn == nil {
			break
		}
		if cn.RemoteAddr().String() != newMain {
			err := fmt.Errorf(
				"sentinel: closing connection to the old main %s",
				cn.RemoteAddr(),
			)
			internal.Logf(err.Error())
			d.pool.Remove(cn, err)
		} else {
			cnsToPut = append(cnsToPut, cn)
		}
	}

	for _, cn := range cnsToPut {
		d.pool.Put(cn)
	}
}

func (d *sentinelFailover) listen(sentinel *sentinelClient) {
	var pubsub *PubSub
	for {
		if pubsub == nil {
			pubsub = sentinel.PubSub()

			if err := pubsub.Subscribe("+switch-main"); err != nil {
				internal.Logf("sentinel: Subscribe failed: %s", err)
				pubsub.Close()
				d.resetSentinel()
				return
			}
		}

		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			internal.Logf("sentinel: ReceiveMessage failed: %s", err)
			pubsub.Close()
			d.resetSentinel()
			return
		}

		switch msg.Channel {
		case "+switch-main":
			parts := strings.Split(msg.Payload, " ")
			if parts[0] != d.mainName {
				internal.Logf("sentinel: ignore new %s addr", parts[0])
				continue
			}

			addr := net.JoinHostPort(parts[3], parts[4])
			internal.Logf(
				"sentinel: new %q addr is %s",
				d.mainName, addr,
			)

			d.closeOldConns(addr)
		}
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
