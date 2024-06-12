package verifyutil

import (
	idvalidator "github.com/guanguans/id-validator"
	"time"
)

// IdInfo 身份证信息
type IdCardInfo struct {
	AddressCode   int      // 身份证归属地
	Address       string   // 地址信息: 省-市-区(县)；
	AddressTree   []string // 省-市-区(县)数组
	Birthday      string   // 生日
	Constellation string   // 星座
	ChineseZodiac string   // 属相
	Sex           string   // 性别;男|女
	IdLen         int      // 身份证长度
	Province      string   // 省
	City          string   // 市
	County        string   // 区
	Age           int      // 年龄
}

/*
* @Description: 获取当前年龄
* @Author: LiuQHui
* @Param year
* @Return int
* @Return error
* @Date 2022-12-01 22:43:14
 */
func GetCurrentAge(year string) (int, error) {
	birthTime, err := time.Parse("2006-01-02", year)
	if err != nil {
		return 0, err
	}
	nowTime := time.Now()

	age := nowTime.Year() - birthTime.Year()

	if nowTime.Month() < birthTime.Month() {
		age--
	} else if nowTime.Month() == birthTime.Month() {
		if nowTime.Day() < birthTime.Day() {
			age--
		}
	}
	return age, nil
}

/*
* @Description: 解析身份证信息
* @Author: LiuQHui
* @Param idCard
* @Return *IdCardInfo
* @Return error
* @Date 2024-06-12 11:10:09
 */
func ParseIdCard(idCard string) (*IdCardInfo, error) {
	cardInfo, err := idvalidator.GetInfo(idCard, false)
	if err != nil {
		return nil, err
	}
	idMsg := IdCardInfo{
		AddressCode:   cardInfo.AddressCode,
		Address:       cardInfo.Address,
		AddressTree:   cardInfo.AddressTree,
		Birthday:      cardInfo.Birthday.Format("2006-01-02"),
		Constellation: cardInfo.Constellation,
		ChineseZodiac: cardInfo.ChineseZodiac,
		IdLen:         cardInfo.Length,
	}
	// 生日
	birthday := cardInfo.Birthday.Format("2006-01-02")
	idMsg.Sex = "男"
	if cardInfo.Sex != 1 {
		idMsg.Sex = "女"
	}
	// 计算当前年龄
	if getAge, err := GetCurrentAge(birthday); err == nil {
		idMsg.Age = getAge
	}
	// 省市区
	if len(cardInfo.AddressTree) >= 3 {
		idMsg.Province = cardInfo.AddressTree[0]
		idMsg.City = cardInfo.AddressTree[1]
		idMsg.County = cardInfo.AddressTree[2]
	}
	return &idMsg, nil
}

/*
* @Description: 身份证是否合法
* @Author: LiuQHui
* @Param card
* @Return bool
* @Date 2024-06-12 11:10:03
 */
func VerifyIdCard(card string) bool {
	return idvalidator.IsValid(card, false)
}
