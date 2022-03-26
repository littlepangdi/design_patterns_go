package main

import (
	"fmt"
	"math/rand"
)

/*
	The Template Method pattern is a behavioral design pattern that defines the framework of an algorithm in a superclass,
		allowing subclasses to override specific steps of the algorithm without modifying the structure.

	1. algorithm interface contains several steps
	2. make concrete algorithm
	3. client owns an algorithm which can be initialized with different instances
	4. client calls algorithm.run()
*/
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//algorithm in steps
type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publish()
}

//complete algorithm
type otp struct {
	iOtp iOtp
}

func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	o.iOtp.publish()
	return nil
}

//concrete algorithm realize 1
type sms struct {
	otp
}

func (s *sms) genRandomOTP(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fmt.Printf("SMS: generating random otp %v\n", b)
	return string(b)
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publish() {
	fmt.Printf("SMS: publishing \n")
}

type email struct {
	otp
}

func (s *email) genRandomOTP(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fmt.Printf("EMAIL: generating random otp %v\n", b)
	return string(b)
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func (s *email) publish() {
	fmt.Printf("EMAIL: publishing \n")
}

func RunTemplate() {
	smsOTP := &sms{}
	emailOTP := &email{}

	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(5)
}
