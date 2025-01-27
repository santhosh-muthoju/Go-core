package main

import (
	attachments "interviewPrep/attachments/functions"
)

func main() {
	sampleAcc := attachments.NewAccount(2000.0)
	attachments.SamplePrintClosure(sampleAcc)
}
