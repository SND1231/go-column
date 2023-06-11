--- テスト用のデータベース作成
CREATE DATABASE unit_test;

--- テストデータ
insert into user (name, age) values 
('taro', 25),
('jiro', 23);