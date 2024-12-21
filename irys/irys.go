package irys

type Client interface {
	Upload(data []byte) (string, error)       //Uploads data to the network and returns the CID
	
	GetUploadPrice(data []byte) (uint, error) //Gets the price to upload data to the network in lamports
	Fund(lamports uint) error                 //Funds the client's wallet with lamports

}
