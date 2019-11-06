package go_testing_tools_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	go_testing_tools "github.com/salmander/go-testing-tools"
	"github.com/salmander/go-testing-tools/gomock_mocks"
)

var _ = Describe("ProductSearch", func() {
	var ctrl *gomock.Controller

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("when no Product is found", func() {
		It("returns an error", func() {
			//Arrange
			ean := "1234"

			mockProductRepository := gomock_mocks.NewMockProductRepository(ctrl)
			mockLogger := gomock_mocks.NewMockCustomLogger(ctrl)

			productRepositoryError := errors.New("some error")

			gomock.InOrder(
				mockProductRepository.EXPECT().
					FindProductByEan(ean).
					Times(1).
					Return(go_testing_tools.Product{}, productRepositoryError),
				mockLogger.EXPECT().
					Log(gomock.Any(), ean, productRepositoryError).Times(1),
			)

			productSearch := go_testing_tools.ProductSearch{
				ProductRepo: mockProductRepository,
				Logger:      mockLogger,
			}

			// Act
			actual, err := productSearch.GetProduct(ean)

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(go_testing_tools.ProductRetrieveError(productRepositoryError)))
			Expect(actual).To(Equal(go_testing_tools.Product{}))
		})
	})
})
