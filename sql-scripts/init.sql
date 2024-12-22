USE jumpStart;
SET NAMES 'utf8mb4';

CREATE TABLE tb_share(
    id INT AUTO_INCREMENT PRIMARY KEY, 
    nameShare VARCHAR(255), 
    dateShare DATE,openShare FLOAT, 
    highShare FLOAT, 
    lowShare FLOAT, 
    closeShare FLOAT, 
    volumeShare FLOAT
);

CREATE TABLE tb_news(
    id INT AUTO_INCREMENT PRIMARY KEY, 
    news TEXT, 
    dateNews DATE,
    datePublished DATE, 
    isApproved BOOLEAN
);


CREATE TABLE tb_investor(
   idInvestor INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
   investorName VARCHAR(255) NOT NULL ,
   investorEmail VARCHAR(255) NOT NULL,
   investorPassword VARCHAR(255) NOT NULL ,
   investorRole VARCHAR(5) NOT NULL CHECK (investorRole IN ('USER', 'ADMIN')),
   validationCode VARCHAR(255) ,
   isAccountValid BOOLEAN NOT NULL
);

CREATE TABLE tb_operationAsset(
 idAsset INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
 assetName VARCHAR(255) NOT NULL ,
 assetType VARCHAR(255) NOT NULL ,
 assetCode VARCHAR(255) NOT NULL,
 assetQuantity DOUBLE NOT NULL CHECK (assetQuantity > 0),
 assetValue DOUBLE NOT NULL CHECK (assetValue > 0),
 operationType VARCHAR(255) NOT NULL CHECK (operationType IN ('BUY', 'SELL')),
 operationDate DATE NOT NULL,
 idInvestor INT NOT NULL,
 isProcessedAlready BOOLEAN
);

CREATE TABLE tb_wallet(
idWallet INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
balance DOUBLE NOT NULL,
idInvestor INT NOT NULL
);

CREATE TABLE tb_walletAsset(
idWalletAsset INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
assetName VARCHAR(255) NOT NULL ,
assetType VARCHAR(255) NOT NULL ,
assetQuantity DOUBLE NOT NULL ,
idWallet INT NOT NULL 
);

CREATE TABLE tb_walletOperation(
idWalletOperation INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
operationType VARCHAR(255) NOT NULL CHECK (operationType IN ('WITHDRAW', 'DEPOSIT')),
operationValue DOUBLE NOT NULL,
operationDate DATE NOT NULL,
idInvestor INT NOT NULL
);

CREATE TABLE list_asset(
    idList INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nameAsset VARCHAR(255) NOT NULL,
    acronymAsset VARCHAR(255) NOT NULL,
    url_image VARCHAR(255),
    typeAsset VARCHAR(255) NOT NULL CHECK (typeAsset IN ('CRYPTO', 'COIN','SHARE'))
);

ALTER TABLE tb_news CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE tb_operationAsset ADD FOREIGN KEY (idInvestor) REFERENCES tb_investor(idInvestor);
ALTER TABLE tb_wallet ADD FOREIGN KEY (idInvestor) REFERENCES tb_investor(idInvestor);
ALTER TABLE tb_walletOperation ADD FOREIGN KEY (idInvestor) REFERENCES tb_investor(idInvestor);
ALTER TABLE tb_walletAsset ADD FOREIGN KEY (idWallet) REFERENCES tb_wallet(idWallet);


-- INSERT INTO tb_news (id,news,dateNews,datePublished,isApproved) VALUES
-- (1,'{"CRYPTO":"Canon apresenta cÔö£├│meras EOS R1 e EOS R5 Mark II como novas mirrorless"}',     '2024-10-11','2024-10-11', 1),
-- (2,'{"CRYPTO":"Novas cÔö£├│maras EOS R1 e EOS R5 Mark II da Canon sÔö£├║o uma ode Ôö£├í tecnologia"}',   '2024-10-11','2024-10-11', 1),
-- (3,'{"CRYPTO":"Tether anuncia fim de cunhagem de USDT nas redes EOS e Algorand"}',           '2024-10-11','2024-10-11', 1),
-- (4,'{"ACAO":"Ibovespa fecha em queda, aos 127.599 pontos, apÔö£Ôöés balanÔö£┬║os e com IPCA morno"}', '2024-10-11','2024-10-11', 1),
-- (5,'{"ACAO":"Magazine Luiza (MGLU3): Resultados sÔö£Ôöélidos no 1T24"}',                          '2024-10-11','2024-10-11', 1);

INSERT INTO tb_share (id, nameShare, dateShare, openShare, highShare, lowShare, closeShare, volumeShare) VALUES 
(71,	'PETR4.SA',	'2024-10-12',	 37.6,  37.65,	37.32, 37.62,	16343000),
(72,	'BBAS3.SA',	'2024-10-12',	 26.28,	26.46,	26.17, 26.33,	12175400),
(73,	'ITSA4.SA',	'2024-10-12',	 10.52,	10.54,	10.44, 10.47,	11660200),
(75,    'VALE3.SA',	'2024-10-12',	 60.99, 62.27,	60.98, 62.13, 20939400),
(76,	'CMIG4.SA',	'2024-10-12',	 11.12,	11.15,	10.93, 11.02,	12321600),
(77,	'SANB11.SA','2024-10-12',    28.74,	28.85,	28.52, 28.58,	4622300),
(78,	'USIM5.SA', '2024-10-12',    6.27,	6.27,	  6.17,	 6.25,	8909300),
(79,	'ABEV3.SA',	'2024-10-12',	 13.09,	13.12,	12.84, 12.88,	31193700),
(80,	'MGLU3.SA',	'2024-10-12',	 9.22,	9.47,	  8.93,	 9.45,	28500100);

INSERT INTO tb_share (id, nameShare, dateShare, openShare, highShare, lowShare, closeShare, volumeShare) VALUES
(91,    'PETR4.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    37.6,  37.65,  37.32, 37.62,   16343000),
(92,    'BBAS3.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    26.28, 26.46,  26.17, 26.33,   12175400),
(93,    'ITSA4.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    10.52, 10.54,  10.44, 10.47,   11660200),
(95,    'VALE3.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    60.99, 62.27,  60.98, 62.13, 20939400),
(96,    'CMIG4.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    11.12, 11.15,  10.93, 11.02,   12321600),
(97,    'SANB11.SA',	DATE_FORMAT(NOW(), '%Y-%m-%d') ,    28.74,     28.85,  28.52, 28.58,   4622300),
(98,    'USIM5.SA', 	DATE_FORMAT(NOW(), '%Y-%m-%d') ,    6.27,      6.27,     6.17,  6.25,  8909300),
(99,    'ABEV3.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    13.09, 13.12,  12.84, 12.88,   31193700),
(100,   'MGLU3.SA',     DATE_FORMAT(NOW(), '%Y-%m-%d') ,    9.22,  9.47,     8.93,  9.45,  28500100);


INSERT INTO tb_share (id, nameShare, dateShare, openShare, highShare, lowShare, closeShare, volumeShare) VALUES
(101, 'PETR4.SA', '2024-11-14', 38.0, 38.4, 37.8, 38.2, 16500000),
(102, 'PETR4.SA', '2024-11-15', 38.3, 38.6, 38.0, 38.5, 16850000),
(103, 'PETR4.SA', '2024-11-16', 38.5, 39.0, 38.4, 38.8, 17000000),
(104, 'PETR4.SA', '2024-11-17', 39.0, 39.5, 38.9, 39.3, 17500000),
(105, 'PETR4.SA', '2024-11-18', 39.2, 39.6, 39.0, 39.4, 18000000),
(106, 'PETR4.SA', '2024-11-19', 39.5, 39.8, 39.2, 39.7, 18250000),
(107, 'PETR4.SA', '2024-11-20', 39.6, 40.0, 39.3, 39.9, 18500000),
(108, 'PETR4.SA', '2024-11-21', 40.0, 40.5, 39.8, 40.2, 19000000),
(109, 'PETR4.SA', '2024-11-22', 40.2, 40.6, 40.0, 40.5, 19200000),
(110, 'PETR4.SA', '2024-11-23', 40.5, 41.0, 40.3, 40.8, 19500000);


INSERT INTO tb_investor(idInvestor, InvestorName,InvestorEmail,InvestorPassword,InvestorRole,validationCode,isAccountValid) VALUES
(1, "investidor nome","Email@investor.com","passwordInvestor","USER","código",FALSE);

INSERT INTO tb_investor(idInvestor, InvestorName,InvestorEmail,InvestorPassword,InvestorRole,validationCode,isAccountValid) VALUES
(3, "investidor2 nome2","Email2@investor.com","passwordInvestor2","USER","código",FALSE);

INSERT INTO tb_investor(idInvestor, InvestorName,InvestorEmail,InvestorPassword,InvestorRole,validationCode,isAccountValid) VALUES
(2, "admin nome","Email@admin.com","passwordAdmin","ADMIN","código",TRUE);

INSERT INTO tb_operationAsset(idAsset,assetName,assetType,assetCode,assetQuantity,assetValue,operationType,operationDate,idInvestor,isProcessedAlready) VALUES
(1, "BITCOIN","CRYPTOMOEDA","BTC-BRL",0.0003,320.890,"BUY","2024-10-27",1,FALSE);

INSERT INTO tb_operationAsset(idAsset,assetName,assetType,assetCode,assetQuantity,assetValue,operationType,operationDate,idInvestor,isProcessedAlready) VALUES
(2, "BBAS3.SA","AÇÃO","BBAS3.SA",2,14.78,"BUY","2024-10-26",1,FALSE);

INSERT INTO tb_operationAsset(idAsset,assetName,assetType,assetCode,assetQuantity,assetValue,operationType,operationDate,idInvestor,isProcessedAlready) VALUES
(3, "DOLAR","MOEDA","USD-BRL",5,5.40,"BUY","2024-10-26",1,FALSE);

INSERT INTO tb_operationAsset(idAsset,assetName,assetType,assetCode,assetQuantity,assetValue,operationType,operationDate,idInvestor,isProcessedAlready)  VALUES
(4, "ITSA4.SA","AÇÃO","ITSA4.SA",5,18.90,"SELL","2024-10-26",2,TRUE);

INSERT INTO tb_wallet(idWallet,balance,idInvestor) VALUES
(1,1000,1);

INSERT INTO tb_wallet(idWallet,balance,idInvestor) VALUES
(2,200,2);

INSERT INTO tb_wallet(idWallet,balance,idInvestor) VALUES
(3,5098.87,3);

INSERT INTO tb_walletOperation(idWalletOperation,operationType,operationValue,operationDate,idInvestor) VALUES
(3,"WITHDRAW",5400,"2024-10-26",3);

INSERT INTO tb_walletAsset(idWalletAsset,assetName,assetType,assetQuantity,idWallet) VALUES
(3,"DOLAR","MOEDA",40,2);

INSERT INTO tb_walletAsset(idWalletAsset,assetName,assetType,assetQuantity,idWallet) VALUES
(4,"ETH","CRIPTOMOEDA",0.03,3);

-- CRYPTO DETAILS --

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Cardano","ADA-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/cardano.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Axie Infinity","AXS-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/axie-infinity.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Bitcoin","BTC-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/bitcoin.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Chiliz","CHZ-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/chiliz.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("EOS","EOS-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/eos.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Ethereum","ETH-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/ethereum.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Chainlink","LINK-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/chainlink.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Litecoin","LTC-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/litecoin.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Tether USDT","USDT-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/tether.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Stellar","XLM-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/stellar.png","CRYPTO");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("XRP","XRP-BRL","https://cdn.investing.com/crypto-logos/20x20/v2/xrp.png","CRYPTO");

-- CRYPTO DETAILS --

-- COINS DETAILS --
INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Boliviano","BOB-BRL","https://paises.ibge.gov.br/img/bandeiras/BO.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Coroa Dinamarquesa","DKK-BRL","https://paises.ibge.gov.br/img/bandeiras/DK.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Coroa Norueguesa","NOK-BRL","https://paises.ibge.gov.br/img/bandeiras/NO.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Coroa Sueca","SEK-BRL","https://paises.ibge.gov.br/img/bandeiras/SE.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Dólar Americano","USD-BRL","https://paises.ibge.gov.br/img/bandeiras/US.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Dólar australiano","AUD-BRL","https://paises.ibge.gov.br/img/bandeiras/AU.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Dólar Canadense","CAD-BRL","https://paises.ibge.gov.br/img/bandeiras/CA.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Dólar Taiuanês","TWD-BRL","https://paises.ibge.gov.br/img/bandeiras/TH.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Euro","EUR-BRL","https://upload.wikimedia.org/wikipedia/commons/thumb/b/b7/Flag_of_Europe.svg/255px-Flag_of_Europe.svg.png","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Franco Suíço","CHF-BRL","https://paises.ibge.gov.br/img/bandeiras/CH.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Guarani Paraguaio","PYG-BRL","https://paises.ibge.gov.br/img/bandeiras/PY.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Iene Japonês","JPY-BRL","https://paises.ibge.gov.br/img/bandeiras/JP.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Peso Argentino","ARS-BRL","https://paises.ibge.gov.br/img/bandeiras/AR.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Peso Chileno","CLP-BRL","https://paises.ibge.gov.br/img/bandeiras/CL.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Peso Mexicano","MXN-BRL","https://paises.ibge.gov.br/img/bandeiras/MX.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Peso Uruguaio","UYU-BRL","https://paises.ibge.gov.br/img/bandeiras/UY.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Rublo Russo","RUB-BRL","https://upload.wikimedia.org/wikipedia/commons/thumb/f/f3/Flag_of_Russia.svg/255px-Flag_of_Russia.svg.png","COIN");


INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Libra Esterlina","GBP-BRL","https://paises.ibge.gov.br/img/bandeiras/GB.gif","COIN");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Yuan Chinês","CNY-BRL","https://paises.ibge.gov.br/img/bandeiras/CN.gif","COIN");

-- COINS DETAILS --

-- SHARE DETAILS --

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("AMBEV S/A ON (ABEV3.SA)","ABEV3.SA","https://investidor10.com.br/storage/companies/66b65af53af6c.jpg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Banco do Brasil (BBAS3.SA)","BBAS3.SA","https://investidor10.com.br/storage/companies/66b65b3de91ca.jpg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Cemig (CMIG4.SA)","CMIG4.SA","https://investidor10.com.br/storage/companies/5ea0b6985411c.jpeg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Itaúsa PN (ITSA4.SA)","ITSA4.SA","https://investidor10.com.br/storage/companies/66b65b5f49225.jpg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Magazine Luiza S.A. (MGLU3.SA)","MGLU3.SA","https://investidor10.com.br/storage/companies/5e9deee481287.jpeg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("PETROBRAS PN (PETR4.SA)","PETR4.SA","https://investidor10.com.br/storage/companies/5e98b684e5df2.jpeg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Banco Santander (SANB11.SA)","SANB11.SA","https://investidor10.com.br/storage/companies/66cc9b5c55343.jpg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("USIMINAS PNA (USIM5.SA)","USIM5.SA","https://investidor10.com.br/storage/companies/66cc984b21a8c.jpg","SHARE");

INSERT INTO list_asset(nameAsset,acronymAsset,url_image,typeAsset) VALUES
("Vale S.A (VALE3.SA)","VALE3.SA","https://investidor10.com.br/storage/companies/5ed732a0242b3.jpeg","SHARE");


-- SHARE DETAILS --