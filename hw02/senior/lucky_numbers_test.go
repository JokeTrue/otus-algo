package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type TestCase struct {
	Input  int
	Output int
}

func GetOrUpdateTestCase(testCase *TestCase, extension string, value int) *TestCase {
	if testCase == nil {
		testCase = &TestCase{}
	}

	if extension == ".in" {
		testCase.Input = value
	} else if extension == ".out" {
		testCase.Output = value
	}

	return testCase
}

func Test_CountLuckyTickets(t *testing.T) {
	files, err := ioutil.ReadDir("cases")
	require.NoError(t, err)

	cases := make(map[string]*TestCase)
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		caseName := strings.TrimSuffix(file.Name(), ext)

		rawValue, fileErr := ioutil.ReadFile("cases/" + file.Name())
		require.NoError(t, fileErr)

		value, convErr := strconv.Atoi(strings.Trim(string(rawValue), "\r\n"))
		require.NoError(t, convErr)

		cases[caseName] = GetOrUpdateTestCase(cases[caseName], ext, value)
	}

	for _, testCase := range cases {
		require.Equal(t, CountLuckyTickets(testCase.Input), testCase.Output)
	}
}
