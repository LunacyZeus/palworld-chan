package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type OptionSettings struct {
	Difficulty                           string  `ini:"Difficulty"`
	DayTimeSpeedRate                     float64 `ini:"DayTimeSpeedRate"`
	NightTimeSpeedRate                   float64 `ini:"NightTimeSpeedRate"`
	ExpRate                              float64 `ini:"ExpRate"`
	PalCaptureRate                       float64 `ini:"PalCaptureRate"`
	PalSpawnNumRate                      float64 `ini:"PalSpawnNumRate"`
	PalDamageRateAttack                  float64 `ini:"PalDamageRateAttack"`
	PalDamageRateDefense                 float64 `ini:"PalDamageRateDefense"`
	PlayerDamageRateAttack               float64 `ini:"PlayerDamageRateAttack"`
	PlayerDamageRateDefense              float64 `ini:"PlayerDamageRateDefense"`
	PlayerStomachDecreaceRate            float64 `ini:"PlayerStomachDecreaceRate"`
	PlayerStaminaDecreaceRate            float64 `ini:"PlayerStaminaDecreaceRate"`
	PlayerAutoHPRegeneRate               float64 `ini:"PlayerAutoHPRegeneRate"`
	PlayerAutoHpRegeneRateInSleep        float64 `ini:"PlayerAutoHpRegeneRateInSleep"`
	PalStomachDecreaceRate               float64 `ini:"PalStomachDecreaceRate"`
	PalStaminaDecreaceRate               float64 `ini:"PalStaminaDecreaceRate"`
	PalAutoHPRegeneRate                  float64 `ini:"PalAutoHPRegeneRate"`
	PalAutoHpRegeneRateInSleep           float64 `ini:"PalAutoHpRegeneRateInSleep"`
	BuildObjectDamageRate                float64 `ini:"BuildObjectDamageRate"`
	BuildObjectDeteriorationDamageRate   float64 `ini:"BuildObjectDeteriorationDamageRate"`
	CollectionDropRate                   float64 `ini:"CollectionDropRate"`
	CollectionObjectHpRate               float64 `ini:"CollectionObjectHpRate"`
	CollectionObjectRespawnSpeedRate     float64 `ini:"CollectionObjectRespawnSpeedRate"`
	EnemyDropItemRate                    float64 `ini:"EnemyDropItemRate"`
	DeathPenalty                         string  `ini:"DeathPenalty"`
	BEnablePlayerToPlayerDamage          bool    `ini:"bEnablePlayerToPlayerDamage"`
	BEnableFriendlyFire                  bool    `ini:"bEnableFriendlyFire"`
	BEnableInvaderEnemy                  bool    `ini:"bEnableInvaderEnemy"`
	BActiveUNKO                          bool    `ini:"bActiveUNKO"`
	BEnableAimAssistPad                  bool    `ini:"bEnableAimAssistPad"`
	BEnableAimAssistKeyboard             bool    `ini:"bEnableAimAssistKeyboard"`
	DropItemMaxNum                       int     `ini:"DropItemMaxNum"`
	DropItemMaxNum_UNKO                  int     `ini:"DropItemMaxNum_UNKO"`
	BaseCampMaxNum                       int     `ini:"BaseCampMaxNum"`
	BaseCampWorkerMaxNum                 int     `ini:"BaseCampWorkerMaxNum"`
	DropItemAliveMaxHours                float64 `ini:"DropItemAliveMaxHours"`
	BAutoResetGuildNoOnlinePlayers       bool    `ini:"bAutoResetGuildNoOnlinePlayers"`
	AutoResetGuildTimeNoOnlinePlayers    float64 `ini:"AutoResetGuildTimeNoOnlinePlayers"`
	GuildPlayerMaxNum                    int     `ini:"GuildPlayerMaxNum"`
	PalEggDefaultHatchingTime            float64 `ini:"PalEggDefaultHatchingTime"`
	WorkSpeedRate                        float64 `ini:"WorkSpeedRate"`
	BIsMultiplay                         bool    `ini:"bIsMultiplay"`
	BIsPvP                               bool    `ini:"bIsPvP"`
	BCanPickupOtherGuildDeathPenaltyDrop bool    `ini:"bCanPickupOtherGuildDeathPenaltyDrop"`
	BEnableNonLoginPenalty               bool    `ini:"bEnableNonLoginPenalty"`
	BEnableFastTravel                    bool    `ini:"bEnableFastTravel"`
	BIsStartLocationSelectByMap          bool    `ini:"bIsStartLocationSelectByMap"`
	BExistPlayerAfterLogout              bool    `ini:"bExistPlayerAfterLogout"`
	BEnableDefenseOtherGuildPlayer       bool    `ini:"bEnableDefenseOtherGuildPlayer"`
	CoopPlayerMaxNum                     int     `ini:"CoopPlayerMaxNum"`
	ServerPlayerMaxNum                   int     `ini:"ServerPlayerMaxNum"`
	ServerName                           string  `ini:"ServerName"`
	ServerDescription                    string  `ini:"ServerDescription"`
	AdminPassword                        string  `ini:"AdminPassword"`
	ServerPassword                       string  `ini:"ServerPassword"`
	PublicPort                           int     `ini:"PublicPort"`
	PublicIP                             string  `ini:"PublicIP"`
	RCONEnabled                          bool    `ini:"RCONEnabled"`
	RCONPort                             int     `ini:"RCONPort"`
	Region                               string  `ini:"Region"`
	BUseAuth                             bool    `ini:"bUseAuth"`
	BanListURL                           string  `ini:"BanListURL"`
}

type GameSettings struct {
	OptionSettings OptionSettings `ini:"OptionSettings"`
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

func uppercaseFirst(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}
	return ""
}

func setField(obj reflect.Value, fieldName string, value string) error {
	field := obj.FieldByName(fieldName)
	if !field.IsValid() {
		// Try to find the field with uppercase first letter
		field = obj.FieldByName(uppercaseFirst(fieldName))
		if !field.IsValid() {
			return fmt.Errorf("Field %s not found", fieldName)
		}
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Float64:
		f, err := parseFloat(value)
		if err != nil {
			return err
		}
		field.SetFloat(f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(b)
	// ... 其他类型的处理
	default:
		return fmt.Errorf("Unsupported field type: %s", field.Kind())
	}

	return nil
}

func printStructFields(settings OptionSettings) {
	settingsValue := reflect.ValueOf(settings)
	settingsType := settingsValue.Type()

	for i := 0; i < settingsValue.NumField(); i++ {
		fieldName := settingsType.Field(i).Name
		fieldValue := settingsValue.Field(i).Interface()

		switch v := fieldValue.(type) {
		case bool:
			fmt.Printf("%s: %t\n", fieldName, v)
		case int:
			fmt.Printf("%s: %d\n", fieldName, v)
		case float64:
			fmt.Printf("%s: %f\n", fieldName, v)
		case string:
			fmt.Printf("%s: %s\n", fieldName, v)
		default:
			fmt.Printf("%s: %v\n", fieldName, fieldValue)
		}
	}
}

func LoadIni() {
	// Your INI configuration string
	iniConfigString := `[Script/Pal.PalGameWorldSettings]
OptionSettings=(Difficulty=None,DayTimeSpeedRate=1.000000,NightTimeSpeedRate=1.000000,ExpRate=1.000000,PalCaptureRate=1.000000,PalSpawnNumRate=1.000000,PalDamageRateAttack=1.000000,PalDamageRateDefense=1.000000,PlayerDamageRateAttack=1.000000,PlayerDamageRateDefense=1.000000,PlayerStomachDecreaceRate=1.000000,PlayerStaminaDecreaceRate=1.000000,PlayerAutoHPRegeneRate=1.000000,PlayerAutoHpRegeneRateInSleep=1.000000,PalStomachDecreaceRate=1.000000,PalStaminaDecreaceRate=1.000000,PalAutoHPRegeneRate=1.000000,PalAutoHpRegeneRateInSleep=1.000000,BuildObjectDamageRate=1.000000,BuildObjectDeteriorationDamageRate=1.000000,CollectionDropRate=1.000000,CollectionObjectHpRate=1.000000,CollectionObjectRespawnSpeedRate=1.000000,EnemyDropItemRate=1.000000,DeathPenalty=All,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=False,DropItemMaxNum=3000,DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampWorkerMaxNum=15,DropItemAliveMaxHours=1.000000,bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=72.000000,GuildPlayerMaxNum=20,PalEggDefaultHatchingTime=1.000000,WorkSpeedRate=1.000000,bIsMultiplay=False,bIsPvP=False,bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=True,bEnableFastTravel=True,bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName="zixing-docker-generated-26979",ServerDescription="Palworld-Dedicated-Server running in Docker by zixing",AdminPassword="123",ServerPassword="",PublicPort=8211,PublicIP="10.0.0.5",RCONEnabled=true,RCONPort=25575,Region="",bUseAuth=True,BanListURL="https://api.palworldgame.com/api/banlist.txt")`

	cfg, err := ini.Load(strings.NewReader(iniConfigString))
	if err != nil {
		fmt.Printf("Failed to load INI config: %v\n", err)
		return
	}

	//var settings OptionSettings
	section := cfg.Section("Script/Pal.PalGameWorldSettings")
	//log.Println(section)

	//keys := section.Keys()

	var settings OptionSettings

	OptionSettingsKey := section.Key("OptionSettings")
	confStr := OptionSettingsKey.String()
	//log.Println(confStr)

	// 使用正则表达式匹配键值对
	re := regexp.MustCompile(`([a-zA-Z_][a-zA-Z0-9_]*)=([^,]+)`)
	matches := re.FindAllStringSubmatch(confStr, -1)
	//log.Println(matches)

	// 使用反射设置结构体字段值
	settingsValue := reflect.ValueOf(&settings).Elem()
	for _, match := range matches {
		key := match[1]
		value := match[2]

		// 根据键名设置字段值
		if err := setField(settingsValue, key, value); err != nil {
			fmt.Printf("Error setting field %s: %v\n", key, err)
			return
		}
	}

	log.Println(settings)

	printStructFields(settings)

}
