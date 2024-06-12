package cryptutil

import "bytes"

/*
* @Description: 补码
* @Author: LiuQHui
* @Param originByte
* @Param blockSize
* @Return []byte
* @Date 2024-06-12 14:00:09
 */
func _pkcs7Padding(originByte []byte, blockSize int) []byte {
	// 计算补码长度
	padding := blockSize - len(originByte)%blockSize
	// 生成补码
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 追加补码
	return append(originByte, padText...)
}

/*
* @Description: 解码
* @Author: LiuQHui
* @Param originDataByte
* @Return []byte
* @Date 2024-06-12 14:00:18
 */
func _pkcs7UNPadding(originDataByte []byte) []byte {
	length := len(originDataByte)
	unpadding := int(originDataByte[length-1])
	return originDataByte[:(length - unpadding)]
}
