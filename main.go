package main

import "test/helpers"

func main() {

	// err := helpers.GeneratePdf("hello.pdf")
    // if err != nil {
    //     panic(err)
    // }

	err := helpers.MailSender()
    if err != nil {
        panic(err)
    }


}

