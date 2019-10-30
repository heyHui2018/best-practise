package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/heyHui2018/log"
	"google.golang.org/grpc/credentials"
)

// server端调用NewServerTLSFromFile,client端调用NewClientTLSFromFile
func TLS() credentials.TransportCredentials {
	c, err := credentials.NewServerTLSFromFile("conf/server.pem", "conf/server.key")
	if err != nil {
		log.Warnf("TLS,credentials.NewServerTLSFromFile error,err = %v", err)
		return nil
	}
	return c
}

// server端tls.Config设置ClientAuth和ClientCAs,client端tls.Config设置ServerName和RootCAs
func CATLS() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("conf/server/server.pem", "conf/server/server.key")
	if err != nil {
		log.Warnf("CATLS,tls.LoadX509KeyPair error,err = %v", err)
		return nil
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("conf/ca.pem")
	if err != nil {
		log.Warnf("CATLS,ioutil.ReadFile error,err = %v", err)
		return nil
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Warnf("CATLS,certPool.AppendCertsFromPEM error")
		return nil
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 设置证书链,允许包含一个或多个
		ClientAuth:   tls.RequireAndVerifyClientCert, // 要求必须校验客户端的证书
		ClientCAs:    certPool,                       // 设置根证书的集合
	})
	return c
}
