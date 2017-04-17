package easyvk

import (
	"encoding/json"
	"fmt"
)

const (
	notifyBit = 1 << iota
	friendsBit
	photosBit
	audioBit
	videoBit
	_
	_
	pagesBit
	_
	_
	statusBit
	notesBit
	messagesBit
	wallBit
	_
	adsBit
	offlineBit
	docsBit
	groupsBit
	notificationsBit
	statsBit
	_
	emailBit
	_
	_
	_
	_
	_
	marketBit
)

// An Account describes a set of methods
// to work with account.
type Account struct {
	vk *VK
}

// GetInfo returns current account info.
func (a *Account) GetInfo(fields string) (Info, error) {
	params := map[string]string{"fields": fields }
	resp, err := a.vk.Request("account.getInfo", params)
	if err != nil {
		return Info{}, err
	}
	var info Info
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return Info{}, err
	}
	return info, nil
}

// GetProfileInfo returns the current account info.
func (a *Account) GetProfileInfo() (ProfileInfo, error) {
	resp, err := a.vk.Request("account.getProfileInfo", nil)
	if err != nil {
		return ProfileInfo{}, err
	}
	var info ProfileInfo
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return ProfileInfo{}, err
	}
	return info, nil
}

// GetCounters returns values of user counters.
func (a *Account) GetCounters(filter string) (Counters, error) {
	params := map[string]string{"filter": filter }
	resp, err := a.vk.Request("account.getCounters", params)
	if err != nil {
		return Counters{}, err
	}
	var counters Counters
	err = json.Unmarshal(resp, &counters)
	if err != nil {
		return Counters{}, err
	}
	return counters, nil
}

// GetAppPermissions returns settings of the user in this application.
func (a *Account) GetAppPermissions(userID uint) (Permissions, error) {
	params := map[string]string{"user_id": fmt.Sprint(userID) }
	resp, err := a.vk.Request("account.getAppPermissions", params)
	if err != nil {
		return Permissions{}, err
	}
	var response response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return Permissions{}, err
	}

	perm := Permissions{}
	permBits := response.Response

	if permBits&notifyBit != 0 {
		perm.Notify = true
	}
	if permBits&friendsBit != 0 {
		perm.Friends = true
	}
	if permBits&photosBit != 0 {
		perm.Photos = true
	}
	if permBits&audioBit != 0 {
		perm.Audio = true
	}
	if permBits&videoBit != 0 {
		perm.Video = true
	}
	if permBits&pagesBit != 0 {
		perm.Pages = true
	}
	if permBits&statusBit != 0 {
		perm.Status = true
	}
	if permBits&notesBit != 0 {
		perm.Notes = true
	}
	if permBits&messagesBit != 0 {
		perm.Messages = true
	}
	if permBits&wallBit != 0 {
		perm.Wall = true
	}
	if permBits&adsBit != 0 {
		perm.Ads = true
	}
	if permBits&offlineBit != 0 {
		perm.Offline = true
	}
	if permBits&docsBit != 0 {
		perm.Docs = true
	}
	if permBits&groupsBit != 0 {
		perm.Groups = true
	}
	if permBits&notificationsBit != 0 {
		perm.Notifications = true
	}
	if permBits&statsBit != 0 {
		perm.Stats = true
	}
	if permBits&emailBit != 0 {
		perm.Email = true
	}
	if permBits&marketBit != 0 {
		perm.Market = true
	}

	return perm, nil
}
