package pipeline

import (
	"io"
	"os"
	"reflect"

	"code.cloudfoundry.org/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/snickers/snickers/db/memory"
	"github.com/snickers/snickers/downloaders"
	"github.com/snickers/snickers/types"
)

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

var _ = Describe("Pipeline", func() {
	var (
		logger *lagertest.TestLogger
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
	})

	Context("Pipeline", func() {
		It("Should get the HTTPDownload function if source is HTTP", func() {
			jobSource := "http://flv.io/KailuaBeach.mp4"
			downloadFunc := GetDownloadFunc(jobSource)
			funcPointer := reflect.ValueOf(downloadFunc).Pointer()
			expected := reflect.ValueOf(downloaders.HTTPDownload).Pointer()
			Expect(funcPointer).To(BeIdenticalTo(expected))
		})

		It("Should get the S3Download function if source is S3", func() {
			jobSource := "http://AWSKEY:AWSSECRET@BUCKET.s3.amazonaws.com/OBJECT"
			downloadFunc := GetDownloadFunc(jobSource)
			funcPointer := reflect.ValueOf(downloadFunc).Pointer()
			expected := reflect.ValueOf(downloaders.S3Download).Pointer()
			Expect(funcPointer).To(BeIdenticalTo(expected))
		})
	})

	Context("when calling Swap Cleaner", func() {
		It("should remove local source and local destination", func() {
			dbInstance, _ := memory.GetDatabase()
			dbInstance.ClearDatabase()
			currentDir, _ := os.Getwd()

			exampleJob := types.Job{
				ID:               "123",
				Source:           "http://source.here.mp4",
				Destination:      "s3://user@pass:/bucket/",
				Preset:           types.Preset{Name: "presetHere", Container: "mp4"},
				Status:           types.JobCreated,
				Details:          "",
				LocalSource:      "/tmp/123/src/KailuaBeach.mp4",
				LocalDestination: "/tmp/123/dst/KailuaBeach.webm",
			}

			dbInstance.StoreJob(exampleJob)

			os.MkdirAll("/tmp/123/src/", 0777)
			os.MkdirAll("/tmp/123/dst/", 0777)

			cp(exampleJob.LocalSource, currentDir+"/../fixtures/videos/nyt.mp4")
			cp(exampleJob.LocalDestination, currentDir+"/../fixtures/videos/nyt.mp4")

			Expect(exampleJob.LocalSource).To(BeAnExistingFile())
			Expect(exampleJob.LocalDestination).To(BeAnExistingFile())

			CleanSwap(dbInstance, exampleJob.ID)

			Expect(exampleJob.LocalSource).To(Not(BeAnExistingFile()))
			Expect(exampleJob.LocalDestination).To(Not(BeAnExistingFile()))
		})
	})
})
