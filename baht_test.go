package bahtgo

import (
	"strconv"
	"testing"
	"time"
)

func TestShouldConvertInStringFormat(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "-11", expected: "ลบสิบเอ็ดบาทถ้วน"},
		{input: "-100", expected: "ลบหนึ่งร้อยบาทถ้วน"},
		{input: "-100.50", expected: "ลบหนึ่งร้อยบาทห้าสิบสตางค์"},
		{input: "-0.50", expected: "ลบห้าสิบสตางค์"},
		{input: "-0", expected: "ศูนย์บาทถ้วน"},
		{input: "0", expected: "ศูนย์บาทถ้วน"},
		{input: "0.50", expected: "ห้าสิบสตางค์"},
		{input: "0.99", expected: "เก้าสิบเก้าสตางค์"},
		{input: "011", expected: "สิบเอ็ดบาทถ้วน"},
		{input: "11", expected: "สิบเอ็ดบาทถ้วน"},
		{input: "21.11", expected: "ยี่สิบเอ็ดบาทสิบเอ็ดสตางค์"},
		{input: "123", expected: "หนึ่งร้อยยี่สิบสามบาทถ้วน"},
		{input: "200.50", expected: "สองร้อยบาทห้าสิบสตางค์"},
		{input: "1234", expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{input: "1234567", expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน"},
		{input: "04123001998830750501", expected: "สี่ล้านหนึ่งแสนสองหมื่นสามพันเอ็ดล้านเก้าแสนเก้าหมื่นแปดพันแปดร้อยสามสิบล้านเจ็ดแสนห้าหมื่นห้าร้อยเอ็ดบาทถ้วน"},
		// {input: "123456789", expected: "ห้าล้านสามพันสี่หมื่นห้าร้อยสิบเจ็ด"},
	}
	timer := time.Now()
	for _, testCase := range testCases {
		actual := convert(testCase.input)
		if actual != testCase.expected {
			t.Errorf(`convert("%s") = "%s", expected "%s"`, testCase.input, actual, testCase.expected)
		}
	}

	output := func() string {
		if time.Since(timer).Milliseconds() > 0 {
			return strconv.FormatInt(time.Since(timer).Milliseconds(), 10) + " ms"
		} else {
			return strconv.FormatInt(time.Since(timer).Microseconds(), 10) + " μs"
		}
	}()

	t.Logf("TestShouldConvertInStringFormat() ran in %s", output)
}

// func TestShouldConvertInFloatFormat(t *testing.T) {
// 	testCases := []struct {
// 		input    float64
// 		expected string
// 	}{
// 		{input: -11, expected: "ลบสิบเอ็ดบาทถ้วน"},
// 		{input: -100, expected: "ลบหนึ่งร้อยบาทถ้วน"},
// 		{input: -100.50, expected: "ลบหนึ่งร้อยบาทห้าสิบสตางค์"},
// 		{input: -0.50, expected: "ลบห้าสิบสตางค์"},
// 		{input: -0, expected: "ศูนย์บาทถ้วน"},
// 		{input: 0, expected: "ศูนย์บาทถ้วน"},
// 		{input: 0.50, expected: "ห้าสิบสตางค์"},
// 		{input: 0.99, expected: "เก้าสิบเก้าสตางค์"},
// 		{input: 11, expected: "สิบเอ็ดบาทถ้วน"},
// 		{input: 21.11, expected: "ยี่สิบเอ็ดบาทสิบเอ็ดสตางค์"},
// 		{input: 123, expected: "หนึ่งร้อยยี่สิบสามบาทถ้วน"},
// 		{input: 200.50, expected: "สองร้อยบาทห้าสิบสตางค์"},
// 		{input: 1234, expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
// 		{input: 1234567, expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทถ้วน"},
// 		{input: 4123001998830750720, expected: "สี่ล้านหนึ่งแสนสองหมื่นสามพันเอ็ดล้านเก้าแสนเก้าหมื่นแปดพันแปดร้อยสามสิบล้านเจ็ดแสนห้าหมื่นเจ็ดร้อยยี่สิบบาทถ้วน"},
// 		// {input: "123456789", expected: "ห้าล้านสามพันสี่หมื่นห้าร้อยสิบเจ็ด"},
// 	}
// 	timer := time.Now()
// 	for _, testCase := range testCases {
// 		actual := convert(testCase.input)
// 		if actual != testCase.expected {
// 			t.Errorf(`convert(%f) = "%s", expected "%s"`, testCase.input, actual, testCase.expected)
// 		}
// 	}

// 	output := func() string {
// 		if time.Since(timer).Milliseconds() > 0 {
// 			return strconv.FormatInt(time.Since(timer).Milliseconds(), 10) + " ms"
// 		} else {
// 			return strconv.FormatInt(time.Since(timer).Microseconds(), 10) + " μs"
// 		}
// 	}()

// 	t.Logf("TestShouldConvertInStringFormat() ran in %s", output)
// }
