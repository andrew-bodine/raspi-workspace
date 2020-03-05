package proxy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pb "github.com/andrew-bodine/release/pkg/protobufs"
	"github.com/andrew-bodine/release/pkg/proxy"
	"github.com/andrew-bodine/release/pkg/proxy/fakes"
)

var _ = Describe("Cacher", func() {
	var (
		lister                *fakes.FakeReleaseLister
		cacher                proxy.Cacher
		zeroValueReleaseCache proxy.ReleaseCache
	)

	Context("with the default raspi release cacher", func() {
		BeforeEach(func() {
			lister = &fakes.FakeReleaseLister{}

			cacher = proxy.NewCacher(lister.Spy)
		})

		Context("that is stopped", func() {
			Context("when Run is called", func() {
				BeforeEach(func() {
					go func() {
						err := cacher.Run()
						Expect(err).NotTo(HaveOccurred())
					}()
				})

				AfterEach(func() {
					err := cacher.Stop()
					Expect(err).NotTo(HaveOccurred())
				})

				It("changes to running state", func() {
					Eventually(func() proxy.CacherState {
						return cacher.State()
					}).Should(Equal(proxy.CacherStateRunning))
				})

				It("starts attempting to build the release cache", func() {
					Eventually(func() int {
						return lister.CallCount()
					}).ShouldNot(Equal(0))
				})
			})

			Context("when Stop is called", func() {
				var err error

				JustBeforeEach(func() {
					err = cacher.Stop()
				})

				It("doesn't return an error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("stays in stopped state", func() {
					Expect(cacher.State()).To(Equal(proxy.CacherStateStopped))
				})

				It("doesn't begin building the release cache", func() {
					Expect(lister.CallCount()).To(Equal(0))
				})
			})

			Context("when Cache is called", func() {
				var releaseCache proxy.ReleaseCache

				JustBeforeEach(func() {
					releaseCache = cacher.Cache()
				})

				It("returns an empty cache", func() {
					Expect(releaseCache).To(Equal(zeroValueReleaseCache))
				})
			})
		})

		Context("that is running and should explicitly stop", func() {
			BeforeEach(func() {
				go func() {
					err := cacher.Run()
					Expect(err).NotTo(HaveOccurred())
				}()

				Eventually(func() proxy.CacherState {
					return cacher.State()
				}).Should(Equal(proxy.CacherStateRunning))
			})

			Context("when Stop is called", func() {
				BeforeEach(func() {
					err := cacher.Stop()
					Expect(err).NotTo(HaveOccurred())
				})

				It("changes to stopped state", func() {
					Eventually(func() proxy.CacherState {
						return cacher.State()
					}).Should(Equal(proxy.CacherStateStopped))
				})

				It("returns an empty cache again", func() {
					Expect(cacher.Cache()).To(Equal(zeroValueReleaseCache))
				})
			})
		})

		Context("that is running", func() {
			BeforeEach(func() {
				go func() {
					err := cacher.Run()
					Expect(err).NotTo(HaveOccurred())
				}()

				Eventually(func() proxy.CacherState {
					return cacher.State()
				}).Should(Equal(proxy.CacherStateRunning))
			})

			AfterEach(func() {
				err := cacher.Stop()
				Expect(err).NotTo(HaveOccurred())
			})

			Context("when Run is called", func() {
				BeforeEach(func() {
					err := cacher.Run()
					Expect(err).NotTo(HaveOccurred())
				})

				It("stays in running state", func() {
					Expect(cacher.State()).To(Equal(proxy.CacherStateRunning))
				})
			})

			Context("when Cache is called", func() {
				var releaseCache proxy.ReleaseCache

				BeforeEach(func() {
					lister.Returns([]pb.Release{{
						Identifier:  "test-release-arm-v0.0.1",
						Description: "This is a test release example.",
						PublishedAt: "now",
						Assets:      []*pb.Release_Asset{},
					}})
				})

				JustBeforeEach(func() {
					releaseCache = cacher.Cache()
				})

				It("returns expected release cache", func() {
					Expect(releaseCache).NotTo(Equal(zeroValueReleaseCache))
				})
			})
		})
	})
})
