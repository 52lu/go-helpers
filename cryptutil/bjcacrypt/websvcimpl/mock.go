package websvcimpl

func getTestResponse(apiName string) string {
	var response string
	switch apiName {
	case ApiGetCert:
		response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
      <GetCertResponse xmlns="http://tempuri.org/" xmlns:ns2="http://cxf.apache.org/transports/http-jetty/configuration" xmlns:ns3="http://cxf.apache.org/configuration/security">
         <GetCert>MIIC6TCCAoygAwIBAgIIITbnrxaPKJgwDAYIKoEcz1UBg3UFADBQMQswCQYDVQQGEwJDTjENMAsGA1UECgwETkhTQTEyMDAGA1UEAwwpUHJpdmF0ZSBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgT2YgTkhTQSBTTTIwHhcNMjMwMzA3MDE1MDE1WhcNNDMwMzAyMDE1MDE1WjCBkDELMAkGA1UEBhMCQ04xDTALBgNVBAoMBE5IU0ExCzAJBgNVBAgMAjExMQswCQYDVQQHDAIwMDELMAkGA1UEBwwCMDAxCzAJBgNVBAsMAjAxMRswGQYDVQQMDBIxMTExMDAwME1CMTUwMDMyMjgxITAfBgNVBAMMGOWMl+S6rOW4guWMu+eWl+S/nemanOWxgDBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABBuJ2h1O5VW2cmin6eEmUlVWmeGLd7CtbPIgfdj0S1qBiEeFEOVmeNJg2ae+ht99YAnS53+61A3wX7wX6d4s4yCjggELMIIBBzALBgNVHQ8EBAMCA8gwgbcGA1UdHwSBrzCBrDAtoCugKYYnaHR0cDovL2NjZW5jYS5uaHNhLmdvdi5jbi9jcmwvY3JsMTQuY3JsMHugeaB3hnVsZGFwOi8vbGNlbmNhLm5oc2EuZ292LmNuL0NOPWNybDE0LE9VPUNSTCxPPU5IU0EsQz1DTj9jZXJ0aWZpY2F0ZVJldm9jYXRpb25MaXN0P2Jhc2U/b2JqZWN0Y2xhc3M9Y1JMRGlzdHJpYnV0aW9uUG9pbnQwHQYDVR0OBBYEFDxPPaPoRXQnB0X7JT8qN7dPYV4ZMB8GA1UdIwQYMBaAFEhEvyjvq9i8FB7d4K9+6EWyKz+WMAwGCCqBHM9VAYN1BQADSQAwRgIhAPwx4bcNoR6jF+JpqULxDUFnZV4uAt2shZI+mYld0c1HAiEAmqOa/YeTt5dtDoE4Zc4ZPMjM/W4hNfoSL33qmtKr5oE=</GetCert>
      </GetCertResponse>
   </soap:Body>
</soap:Envelope>
`
	case ApiEncodeEnvelopedData:
		response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
      <EncodeEnvelopedDataResponse xmlns="http://tempuri.org/" xmlns:ns2="http://cxf.apache.org/transports/http-jetty/configuration" xmlns:ns3="http://cxf.apache.org/configuration/security">
         <EncodeEnvelopedData>QXV0aG9yaXR5IE9mIE5IU0dA0l8O</EncodeEnvelopedData>
      </EncodeEnvelopedDataResponse>
   </soap:Body>
</soap:Envelope>`
	case ApiVerifySignedDataByP7Attach:
		response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
      <VerifySignedDataByP7AttachResponse xmlns="http://tempuri.org/" xmlns:ns2="http://cxf.apache.org/transports/http-jetty/configuration" xmlns:ns3="http://cxf.apache.org/configuration/security">
         <VerifySignedDataByP7Attach>true</VerifySignedDataByP7Attach>
         <IsTest>true</IsTest>
      </VerifySignedDataByP7AttachResponse>
   </soap:Body>
</soap:Envelope>`

	case ApiDecodeEnvelopedData:
		response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
      <DecodeEnvelopedDataResponse xmlns="http://tempuri.org/" xmlns:ns2="http://cxf.apache.org/transports/http-jetty/configuration" xmlns:ns3="http://cxf.apache.org/configuration/security">
         <DecodeEnvelopedData>解密内容112323232</DecodeEnvelopedData>
         <IsTest>true</IsTest>
      </DecodeEnvelopedDataResponse>
   </soap:Body>
</soap:Envelope>`
	case ApiSignDataByP7Attach:
		response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
      <SignDataByP7AttachResponse xmlns="http://tempuri.org/" xmlns:ns2="http://cxf.apache.org/transports/http-jetty/configuration" xmlns:ns3="http://cxf.apache.org/configuration/security">
         <SignDataByP7AttachResult>MIIHUwYKKoEcz1UGAQQCAqCCB0Mwggc/AgEBMQ8wDQYJKoEcz1UBgxECBQAwggRzP</SignDataByP7AttachResult>
      </SignDataByP7AttachResponse>
   </soap:Body>
</soap:Envelope>`

	}
	return response
}
