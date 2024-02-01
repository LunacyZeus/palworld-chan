package setting

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/utils"
	"reflect"
	"regexp"
)

func ParseGameWorldSettingsFromString(settingsString string) (settings OptionSettings) {
	// 使用正则表达式匹配键值对
	re := regexp.MustCompile(`([a-zA-Z_][a-zA-Z0-9_]*)=([^,]+)`)
	matches := re.FindAllStringSubmatch(settingsString, -1)
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
	return settings
}

func LoadGameWorldSettings(iniPath string) (err error) {
	// 加载INI文件
	cfg, err := ini.Load(iniPath)
	if err != nil {
		return
	}
	//var settingsString string

	//获取配置项
	sectionName := "/Script/Pal.PalGameWorldSettings"
	section, err := cfg.GetSection(sectionName)
	if err != nil {
		//settingsString = "(Difficulty=None,DayTimeSpeedRate=1.000000,NightTimeSpeedRate=1.000000,ExpRate=1.000000,PalCaptureRate=1.000000,PalSpawnNumRate=1.000000,PalDamageRateAttack=1.000000,PalDamageRateDefense=1.000000,PlayerDamageRateAttack=1.000000,PlayerDamageRateDefense=1.000000,PlayerStomachDecreaceRate=1.000000,PlayerStaminaDecreaceRate=1.000000,PlayerAutoHPRegeneRate=1.000000,PlayerAutoHpRegeneRateInSleep=1.000000,PalStomachDecreaceRate=1.000000,PalStaminaDecreaceRate=1.000000,PalAutoHPRegeneRate=1.000000,PalAutoHpRegeneRateInSleep=1.000000,BuildObjectDamageRate=1.000000,BuildObjectDeteriorationDamageRate=1.000000,CollectionDropRate=1.000000,CollectionObjectHpRate=1.000000,CollectionObjectRespawnSpeedRate=1.000000,EnemyDropItemRate=1.000000,DeathPenalty=All,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=False,DropItemMaxNum=3000,DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampWorkerMaxNum=15,DropItemAliveMaxHours=1.000000,bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=72.000000,GuildPlayerMaxNum=20,PalEggDefaultHatchingTime=72.000000,WorkSpeedRate=1.000000,bIsMultiplay=False,bIsPvP=False,bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=True,bEnableFastTravel=True,bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName=\"palgo\",ServerDescription=\"https://github.com/Hoshinonyaruko/palworld-go\",AdminPassword=\"useradmin\",ServerPassword=\"\",PublicPort=8211,PublicIP=\"\",RCONEnabled=True,RCONPort=25575,Region=\"\",bUseAuth=True,BanListURL=\"https://api.palworldgame.com/api/banlist.txt\")"
		// 解析设置字符串
		//return parseSettings(settingsString), nil
		logger.Debug("解析配置文件失败")
		err = errors.New("解析配置文件失败")
		return
	}

	// 获取OptionSettings项的值
	optionSettingsKey, err := section.GetKey("OptionSettings")
	if err != nil {
		logger.Debug("解析配置文件失败")
		err = errors.New("解析配置文件失败")
		return
	}
	settingsString := optionSettingsKey.String()
	fmt.Printf("setting:%s\n", settingsString)

	optionSettings := ParseGameWorldSettingsFromString(settingsString)
	//printStructFields(optionSettings)
	settingJson, err := utils.ToJSONString(optionSettings)
	if err != nil {
		logger.Debug("转换json失败")
		return
	}
	fmt.Printf("settingJson:%s\n", settingJson)

	return
}
