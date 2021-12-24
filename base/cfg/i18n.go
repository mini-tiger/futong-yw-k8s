package cfg

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundleI18n *i18n.Bundle

func InitI18n(path string, filenameSlice []string) {
	// Simplified Chinese is read by default
	bundleI18n = i18n.NewBundle(language.SimplifiedChinese)
	bundleI18n.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// Load language configuration file
	for _, filename := range filenameSlice {
		file := fmt.Sprintf("%s/%s", path, filename)
		bundleI18n.MustLoadMessageFile(file)
	}
}

func GetI18nMsg(lang string, messageID string) string {
	localizer := i18n.NewLocalizer(bundleI18n, lang)
	message, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
	if err != nil {
		err = fmt.Errorf(
			"failed to get message with lang is %s and message ID is %s, err is: %s",
			lang, messageID, err.Error(),
		)
		Mlog.Error(err.Error())
		return ""
	}
	return message
}
