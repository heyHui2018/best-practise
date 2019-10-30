## 证书制作

### 1、私钥
* openssl genrsa -out server.key 2048

生成RSA私钥，最后一个参数用于指定生成密钥的位数，默认512

* openssl ecparam -genkey -name secp384r1 -out server.key

生成ECC私钥，命令为椭圆曲线密钥参数生成及操作，secp384r1

### 2、自签名公钥
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期，输入命令后会弹出拥有者信息输入框，如下所示：
```
Country Name (2 letter code) [XX]:
State or Province Name (full name) []:
Locality Name (eg, city) [Default City]:
Organization Name (eg, company) [Default Company Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (eg, your name or your server's hostname) []:grpc server name
Email Address []:
```
输入完成后会生成server.key和server.pem两个文件，放入conf文件夹中即可

### 3、应用
* 引入包"google.golang.org/grpc/credentials"
* server端 credentials.NewServerTLSFromFile("conf/server.pem", "conf/server.key")
* client端 credentials.NewClientTLSFromFile("conf/server.pem", "best-practise")

### 4、CA
* 公钥 openssl genrsa -out ca.key 2048
* 私钥 openssl req -new -x509 -days 7200 -key ca.key -out ca.pem
* Server端
    * 生成CSR openssl req -new -key server.key -out server.csr
    * 签发    openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
* Client端
    * 生成Key openssl ecparam -genkey -name secp384r1 -out client.key
    * 生成CSR openssl req -new -key client.key -out client.csr
    * 签发    openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
