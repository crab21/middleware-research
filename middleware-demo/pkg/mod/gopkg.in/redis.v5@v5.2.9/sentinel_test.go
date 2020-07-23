package redis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/redis.v5"
)

var _ = Describe("Sentinel", func() {
	var client *redis.Client

	BeforeEach(func() {
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MainName:    sentinelName,
			SentinelAddrs: []string{":" + sentinelPort},
		})
		Expect(client.FlushDb().Err()).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(client.Close()).NotTo(HaveOccurred())
	})

	It("should facilitate failover", func() {
		// Set value on main, verify
		err := client.Set("foo", "main", 0).Err()
		Expect(err).NotTo(HaveOccurred())

		val, err := sentinelMain.Get("foo").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal("main"))

		// Wait until replicated
		Eventually(func() string {
			return sentinelSubordinate1.Get("foo").Val()
		}, "1s", "100ms").Should(Equal("main"))
		Eventually(func() string {
			return sentinelSubordinate2.Get("foo").Val()
		}, "1s", "100ms").Should(Equal("main"))

		// Wait until subordinates are picked up by sentinel.
		Eventually(func() string {
			return sentinel.Info().Val()
		}, "10s", "100ms").Should(ContainSubstring("subordinates=2"))

		// Kill main.
		sentinelMain.Shutdown()
		Eventually(func() error {
			return sentinelMain.Ping().Err()
		}, "5s", "100ms").Should(HaveOccurred())

		// Wait for Redis sentinel to elect new main.
		Eventually(func() string {
			return sentinelSubordinate1.Info().Val() + sentinelSubordinate2.Info().Val()
		}, "30s", "1s").Should(ContainSubstring("role:main"))

		// Check that client picked up new main.
		Eventually(func() error {
			return client.Get("foo").Err()
		}, "5s", "100ms").ShouldNot(HaveOccurred())
	})

	It("supports DB selection", func() {
		Expect(client.Close()).NotTo(HaveOccurred())

		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MainName:    sentinelName,
			SentinelAddrs: []string{":" + sentinelPort},
			DB:            1,
		})
		err := client.Ping().Err()
		Expect(err).NotTo(HaveOccurred())
	})
})
