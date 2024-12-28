SET NAMES 'utf8mb4';

CREATE TABLE tb_share  (id INT AUTO_INCREMENT PRIMARY KEY, nameShare VARCHAR(255), dateShare DATE,openShare FLOAT, highShare FLOAT, lowShare FLOAT, closeShare FLOAT, volumeShare FLOAT);

CREATE TABLE tb_news(
    id INT AUTO_INCREMENT PRIMARY KEY, 
    news TEXT, 
    dateNews DATE,
    datePublished DATE, 
    isApproved BOOLEAN
);