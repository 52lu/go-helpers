package viperimpl

import "time"

func (v *viperParseInstance) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v *viperParseInstance) GetString(key string) string {
	return v.viper.GetString(key)
}

func (v *viperParseInstance) GetBool(key string) bool {
	return v.viper.GetBool(key)
}

func (v *viperParseInstance) GetInt64(key string) int64 {
	return v.viper.GetInt64(key)
}

func (v *viperParseInstance) GetInt(key string) int {
	return v.viper.GetInt(key)
}

func (v *viperParseInstance) GetFloat64(key string) float64 {
	return v.viper.GetFloat64(key)
}

func (v *viperParseInstance) GetTime(key string) time.Time {
	return v.viper.GetTime(key)
}

func (v *viperParseInstance) GetIntSlice(key string) []int {
	return v.viper.GetIntSlice(key)
}

func (v *viperParseInstance) GetStringSlice(key string) []string {
	return v.viper.GetStringSlice(key)
}

func (v *viperParseInstance) GetStringMap(key string) map[string]interface{} {
	return v.viper.GetStringMap(key)
}

func (v *viperParseInstance) GetStringMapString(key string) map[string]string {
	return v.viper.GetStringMapString(key)
}

func (v *viperParseInstance) GetStringMapStringSlice(key string) map[string][]string {
	return v.viper.GetStringMapStringSlice(key)
}
