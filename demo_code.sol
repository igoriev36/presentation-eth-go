
//start REST OMIT

http://somesite/sendvalue?from=dave&to=SauSheong&amount=42.00

//end REST OMIT

//start RESTx OMIT

http://somesite/sendvalue?from=dave&to=SauSheong&amount=42.00&signature=1811c4c90b7b

//end RESTx OMIT


//start skeleton OMIT
//end skeleton OMIT


	ownerTx, err = bind.NewTransactor(strings.NewReader(mainKey), ",3ePVWY2LnBEZYD")
	if err != nil {
		log.Printf("Failed to create authorized transactor: %v", err)
		return false
	}
