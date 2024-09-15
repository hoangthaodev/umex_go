package service

import (
	"fmt"

	"umex.com/internal/repo"
	"umex.com/internal/utils/random"
	"umex.com/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	//...
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail

	// 5. check OTP is available

	// 6. user spam...

	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is ::%d\n", otp)

	// 3. save OTP in Redis with expiration time

	// 4. send OTP to email

	return response.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
