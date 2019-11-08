// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BasicTournamentABI is the input ABI used to generate the binding from.
const BasicTournamentABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cooAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"duelResolver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"powerScale_\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"tournamentStartBlock_\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"admissionDuration_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"revivalDuration_\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"ascensionDuration_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"fightDuration_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"cullingDuration_\",\"type\":\"uint24\"},{\"internalType\":\"uint88\",\"name\":\"blueMoldBasePower_\",\"type\":\"uint88\"},{\"internalType\":\"uint24\",\"name\":\"sessionsBetweenMoldDoubling_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"duelTimeoutBlocks_\",\"type\":\"uint24\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ascendingWizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengingWizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"AscensionChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"AscensionComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"}],\"name\":\"AscensionPairUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"AscensionStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCeo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"CEOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCfo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"CFOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCoo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"COOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"duelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"moveSet1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"moveSet2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power2\",\"type\":\"uint256\"}],\"name\":\"DuelEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"duelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAscensionBattle\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commit1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commit2\",\"type\":\"bytes32\"}],\"name\":\"DuelStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"duelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power2\",\"type\":\"uint256\"}],\"name\":\"DuelTimeOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"committingWizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otherWizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"committingWizardNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otherWizardNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"OneSidedCommitAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"OneSidedCommitCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"duelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"committingWizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"otherWizardId\",\"type\":\"uint256\"}],\"name\":\"OneSidedRevealAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseEndedBlock\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sendingWizId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivingWizId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountTransferred\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"PowerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimingWinnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"Revive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"WizardElimination\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"GATE_KEEPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WIZARD_GUILD\",\"outputs\":[{\"internalType\":\"contractWizardGuildInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"acceptAscensionChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"cancelCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"challengeAscending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimingWinnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"allWinners\",\"type\":\"uint256[]\"}],\"name\":\"claimSharedWinnings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimingWinnerId\",\"type\":\"uint256\"}],\"name\":\"claimTheBigCheeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"completeAscension\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"moldyWizardIds\",\"type\":\"uint256[]\"}],\"name\":\"cullMoldedWithMolded\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"survivor\",\"type\":\"uint256\"}],\"name\":\"cullMoldedWithSurvivor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"}],\"name\":\"cullTiredWizards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commit1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commit2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig2\",\"type\":\"bytes\"}],\"name\":\"doubleCommit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"duelId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commit1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commit2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"moveSet1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"moveSet2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt2\",\"type\":\"bytes32\"}],\"name\":\"doubleReveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"duelResolver\",\"outputs\":[{\"internalType\":\"contractDuelResolverInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint88[]\",\"name\":\"powers\",\"type\":\"uint88[]\"}],\"name\":\"enterWizards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAscendingWizardId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlueMoldParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainingWizards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimeParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tournamentStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pauseEndedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"admissionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revivalDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duelTimeoutDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ascensionWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ascensionWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fightWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fightWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolutionWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolutionWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cullingWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cullingWindowDuration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"getWizard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"affinity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"currentDuel\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"ascending\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"ascensionOpponent\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"molded\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sendingWizardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivingWizardId\",\"type\":\"uint256\"}],\"name\":\"giftPower\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"isReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"committingWizardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"otherWizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"oneSidedCommit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"committingWizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"moveSet\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"otherWizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"otherCommit\",\"type\":\"bytes32\"}],\"name\":\"oneSidedReveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseDuration\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"powerScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"resolveOneSidedAscensionBattle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"}],\"name\":\"resolveTimedOutDuel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"revive\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"setCeo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"setCfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"setCoo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"startAscension\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"updateAffinity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"wizardFingerprint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BasicTournamentFuncSigs maps the 4-byte function signature to its string representation.
var BasicTournamentFuncSigs = map[string]string{
	"128183ba": "GATE_KEEPER()",
	"c9eb068b": "WIZARD_GUILD()",
	"9a963fcd": "acceptAscensionChallenge(bytes32)",
	"840a1ff4": "cancelCommitment(uint256)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"fb3790c5": "challengeAscending(uint256,bytes32)",
	"cc4db960": "claimSharedWinnings(uint256,uint256[])",
	"da0cb2ae": "claimTheBigCheeze(uint256)",
	"aedb27fc": "completeAscension()",
	"b047fb50": "cooAddress()",
	"0af29b96": "cullMoldedWithMolded(uint256[])",
	"cbe6549e": "cullMoldedWithSurvivor(uint256[],uint256)",
	"43d9922f": "cullTiredWizards(uint256[])",
	"83197ef0": "destroy()",
	"c51b58aa": "doubleCommit(uint256,uint256,bytes32,bytes32,bytes,bytes)",
	"7aa66a7c": "doubleReveal(uint256,uint256,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
	"810306b9": "duelResolver()",
	"b9d95abb": "enterWizards(uint256[],uint88[])",
	"bfd31c19": "getAscendingWizardId()",
	"079cfa79": "getBlueMoldParameters()",
	"9ce0c954": "getRemainingWizards()",
	"92420c90": "getTimeParameters()",
	"fac8eafc": "getWizard(uint256)",
	"170f8cc8": "giftPower(uint256,uint256)",
	"22f3e2d4": "isActive()",
	"b187bd26": "isPaused()",
	"50df8f71": "isReady(uint256)",
	"01d1c810": "oneSidedCommit(uint256,uint256,bytes32)",
	"cef9b488": "oneSidedReveal(uint256,bytes32,bytes32,bytes32,uint256,bytes32)",
	"136439dd": "pause(uint256)",
	"ad81e4d6": "powerScale()",
	"35a966f1": "resolveOneSidedAscensionBattle(uint256)",
	"feb62755": "resolveTimedOutDuel(uint256,uint256)",
	"8baecc21": "revive(uint256)",
	"88975198": "setCeo(address)",
	"2d46ed56": "setCfo(address)",
	"9986a0c6": "setCoo(address)",
	"58042deb": "startAscension(uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"5a453d40": "updateAffinity(uint256)",
	"3f976ca9": "wizardFingerprint(uint256)",
}

// BasicTournamentBin is the compiled bytecode used for deploying new contracts.
var BasicTournamentBin = "0x60806040523480156200001157600080fd5b5060405162005e9938038062005e9983398181016040526101808110156200003857600080fd5b508051602082015160408301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b0151610160909b0151999a9899979896979596949593949293919290918b898989898989878a8a896000620000ac336001600160e01b0362000b9216565b620000c0826001600160e01b0362000bfc16565b6001600160a01b03811615620000e457620000e4816001600160e01b0362000cf416565b5050438964ffffffffff16116200015c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e76616c69642073746172742074696d650000000000000000000000000000604482015290519081900360640190fd5b60148362ffffff161015620001d257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f54696d656f757420746f6f2073686f7274000000000000000000000000000000604482015290519081900360640190fd5b8062ffffff16826001600160581b03168562ffffff168762ffffff168962ffffff168b63ffffffff168d63ffffffff160202020202026000141562000263576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018062005e766023913960400191505060405180910390fd5b8262ffffff166002028562ffffff161015620002e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f46696768742077696e646f7720746f6f2073686f727400000000000000000000604482015290519081900360640190fd5b8262ffffff168462ffffff1610156200035a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43756c6c696e672077696e646f7720746f6f2073686f72740000000000000000604482015290519081900360640190fd5b62ffffff86860184018501168063ffffffff8916816200037657fe5b0663ffffffff16600014620003ec57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265766976616c2f53657373696f6e206c656e677468206d69736d6174636800604482015290519081900360640190fd5b6040518060a001604052808b64ffffffffff1665ffffffffffff168152602001600065ffffffffffff1681526020018a63ffffffff1681526020018963ffffffff1681526020018562ffffff1663ffffffff16815250600360008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548165ffffffffffff021916908365ffffffffffff160217905550604082015181600001600c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160000160146101000a81548163ffffffff021916908363ffffffff16021790555090505060008963ffffffff168b64ffffffffff1601905060405180608001604052808a63ffffffff16830165ffffffffffff168152602001600065ffffffffffff1681526020018363ffffffff1681526020018962ffffff1663ffffffff16815250600460008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548165ffffffffffff021916908365ffffffffffff160217905550604082015181600001600c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160106101000a81548163ffffffff021916908363ffffffff16021790555090505060405180608001604052808962ffffff16830165ffffffffffff168152602001600065ffffffffffff1681526020018363ffffffff1681526020018862ffffff1663ffffffff16815250600560008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548165ffffffffffff021916908365ffffffffffff160217905550604082015181600001600c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160106101000a81548163ffffffff021916908363ffffffff16021790555090505060405180608001604052808862ffffff168a62ffffff1684010165ffffffffffff168152602001600065ffffffffffff1681526020018363ffffffff1681526020018662ffffff1663ffffffff16815250600660008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548165ffffffffffff021916908365ffffffffffff160217905550604082015181600001600c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160106101000a81548163ffffffff021916908363ffffffff16021790555090505060405180608001604052808662ffffff168962ffffff168b62ffffff168d63ffffffff16860101010165ffffffffffff168152602001600065ffffffffffff1681526020018363ffffffff1681526020018762ffffff1663ffffffff16815250600760008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548165ffffffffffff021916908365ffffffffffff160217905550604082015181600001600c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160106101000a81548163ffffffff021916908363ffffffff16021790555090505060405180608001604052808a63ffffffff16830165ffffffffffff1681526020018363ffffffff168152602001838562ffffff160263ffffffff168152602001856001600160581b0316815250600860008201518160000160006101000a81548165ffffffffffff021916908365ffffffffffff16021790555060208201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600e6101000a8154816001600160581b0302191690836001600160581b031602179055509050505050505050505050505050508a600a60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006001600160a01b03168b6001600160a01b03161415801562000b0e5750600a54604080517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f41fc4f1e00000000000000000000000000000000000000000000000000000000600482015290516001600160a01b03909216916301ffc9a791602480820192602092909190829003018186803b15801562000adf57600080fd5b505afa15801562000af4573d6000803e3d6000fd5b505050506040513d602081101562000b0b57600080fd5b50515b62000b7a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e76616c6964204475656c5265736f6c766572000000000000000000000000604482015290519081900360640190fd5b5050506009969096555062000e829650505050505050565b600054604080516001600160a01b039283168152918316602083015280517f9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc69281900390910190a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331462000c7657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f4f6e6c792043454f000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b62000c8a816001600160e01b0362000dec16565b600154604080516001600160a01b039283168152918316602083015280517f1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe11849281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331462000d6e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f4f6e6c792043454f000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b62000d82816001600160e01b0362000dec16565b600254604080516001600160a01b039283168152918316602083015280517fe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d69281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381161580159062000e1357506000546001600160a01b03828116911614155b62000e7f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e76616c69642043454f206164647265737300000000000000000000000000604482015290519081900360640190fd5b50565b614fe48062000e926000396000f3fe6080604052600436106102515760003560e01c80638897519811610139578063b9d95abb116100b6578063cc4db9601161007a578063cc4db96014610a62578063cef9b48814610ae4578063da0cb2ae14610b2c578063fac8eafc14610b56578063fb3790c514610bcd578063feb6275514610bfd57610251565b8063b9d95abb1461081c578063bfd31c19146108da578063c51b58aa146108ef578063c9eb068b146109d2578063cbe6549e146109e757610251565b80639ce0c954116100fd5780639ce0c954146107b3578063ad81e4d6146107c8578063aedb27fc146107dd578063b047fb50146107f2578063b187bd261461080757610251565b8063889751981461068c5780638baecc21146106bf57806392420c90146106dc5780639986a0c6146107565780639a963fcd1461078957610251565b80632d46ed56116101d257806358042deb1161019657806358042deb1461058f5780635a453d40146105b95780637aa66a7c146105e3578063810306b91461063857806383197ef01461064d578063840a1ff41461066257610251565b80632d46ed561461045157806335a966f1146104845780633f976ca9146104ae57806343d9922f146104ea57806350df8f711461056557610251565b80630af29b96116102195780630af29b9614610352578063128183ba146103cd578063136439dd146103e2578063170f8cc81461040c57806322f3e2d41461043c57610251565b806301d1c8101461025357806301ffc9a7146102895780630519ce79146102d1578063079cfa79146103025780630a0f81681461033d575b005b34801561025f57600080fd5b506102516004803603606081101561027657600080fd5b5080359060208101359060400135610c2d565b34801561029557600080fd5b506102bd600480360360208110156102ac57600080fd5b50356001600160e01b031916611097565b604080519115158252519081900360200190f35b3480156102dd57600080fd5b506102e66110d0565b604080516001600160a01b039092168252519081900360200190f35b34801561030e57600080fd5b506103176110df565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34801561034957600080fd5b506102e661111b565b34801561035e57600080fd5b506102516004803603602081101561037557600080fd5b810190602081018135600160201b81111561038f57600080fd5b8201836020820111156103a157600080fd5b803590602001918460208302840111600160201b831117156103c257600080fd5b50909250905061112a565b3480156103d957600080fd5b506102e6611364565b3480156103ee57600080fd5b506102516004803603602081101561040557600080fd5b503561137c565b34801561041857600080fd5b506102516004803603604081101561042f57600080fd5b5080359060200135611590565b34801561044857600080fd5b506102bd6115ed565b34801561045d57600080fd5b506102516004803603602081101561047457600080fd5b50356001600160a01b031661162d565b34801561049057600080fd5b50610251600480360360208110156104a757600080fd5b50356116ea565b3480156104ba57600080fd5b506104d8600480360360208110156104d157600080fd5b50356117ef565b60408051918252519081900360200190f35b3480156104f657600080fd5b506102516004803603602081101561050d57600080fd5b810190602081018135600160201b81111561052757600080fd5b82018360208201111561053957600080fd5b803590602001918460208302840111600160201b8311171561055a57600080fd5b5090925090506118bf565b34801561057157600080fd5b506102bd6004803603602081101561058857600080fd5b503561194d565b34801561059b57600080fd5b50610251600480360360208110156105b257600080fd5b50356119d9565b3480156105c557600080fd5b50610251600480360360208110156105dc57600080fd5b5035611bc9565b3480156105ef57600080fd5b50610251600480360361010081101561060757600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c08101359060e00135611cee565b34801561064457600080fd5b506102e6611e88565b34801561065957600080fd5b50610251611e97565b34801561066e57600080fd5b506102516004803603602081101561068557600080fd5b5035611ef0565b34801561069857600080fd5b50610251600480360360208110156106af57600080fd5b50356001600160a01b0316611fbb565b610251600480360360208110156106d557600080fd5b503561201a565b3480156106e857600080fd5b506106f161220d565b604080519d8e5260208e019c909c528c8c019a909a5260608c019890985260808b019690965260a08a019490945260c089019290925260e088015261010087015261012086015261014085015261016084015261018083015251908190036101a00190f35b34801561076257600080fd5b506102516004803603602081101561077957600080fd5b50356001600160a01b0316612284565b34801561079557600080fd5b50610251600480360360208110156107ac57600080fd5b5035612341565b3480156107bf57600080fd5b506104d86123f3565b3480156107d457600080fd5b506104d86123f9565b3480156107e957600080fd5b506102516123ff565b3480156107fe57600080fd5b506102e6612516565b34801561081357600080fd5b506102bd612525565b6102516004803603604081101561083257600080fd5b810190602081018135600160201b81111561084c57600080fd5b82018360208201111561085e57600080fd5b803590602001918460208302840111600160201b8311171561087f57600080fd5b919390929091602081019035600160201b81111561089c57600080fd5b8201836020820111156108ae57600080fd5b803590602001918460208302840111600160201b831117156108cf57600080fd5b50909250905061253c565b3480156108e657600080fd5b506104d86128d5565b3480156108fb57600080fd5b506104d8600480360360c081101561091257600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b81111561094457600080fd5b82018360208201111561095657600080fd5b803590602001918460018302840111600160201b8311171561097757600080fd5b919390929091602081019035600160201b81111561099457600080fd5b8201836020820111156109a657600080fd5b803590602001918460018302840111600160201b831117156109c757600080fd5b5090925090506128db565b3480156109de57600080fd5b506102e6612cd1565b3480156109f357600080fd5b5061025160048036036040811015610a0a57600080fd5b810190602081018135600160201b811115610a2457600080fd5b820183602082011115610a3657600080fd5b803590602001918460208302840111600160201b83111715610a5757600080fd5b919350915035612ce9565b348015610a6e57600080fd5b5061025160048036036040811015610a8557600080fd5b81359190810190604081016020820135600160201b811115610aa657600080fd5b820183602082011115610ab857600080fd5b803590602001918460208302840111600160201b83111715610ad957600080fd5b509092509050612df5565b348015610af057600080fd5b50610251600480360360c0811015610b0757600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561312d565b348015610b3857600080fd5b5061025160048036036020811015610b4f57600080fd5b503561360a565b348015610b6257600080fd5b50610b8060048036036020811015610b7957600080fd5b50356136f2565b60408051998a5260208a01989098528888019690965260608801949094526080870192909252151560a086015260c0850152151560e0840152151561010083015251908190036101200190f35b348015610bd957600080fd5b5061025160048036036040811015610bf057600080fd5b508035906020013561386c565b348015610c0957600080fd5b5061025160048036036040811015610c2057600080fd5b5080359060200135613a4d565b610c35613c1f565b82610c3f81613cc5565b82610c4981613e02565b610c538585613e6a565b600085815260106020526040812054151580610c7c575060008581526010602052604090205415155b15610d005760008681526010602052604090205485148015610cab575060008581526010602052604090205486145b610cfc576040805162461bcd60e51b815260206004820152601d60248201527f4d757374207265736f6c766520417363656e73696f6e20426174746c65000000604482015290519081900360640190fd5b5060015b610d08614f1c565b506000868152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff166060830152600101546080820152610d76614f1c565b506000868152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff166060830152600101546080820152600f548814801590610df057506011548814155b8015610e0d575081516001600160581b0316610e0a613f69565b11155b8015610e1f5750606082015160ff1615155b8015610e2d57506080820151155b610e71576040805162461bcd60e51b815260206004820152601060248201526f57697a617264206e6f7420726561647960801b604482015290519081900360640190fd5b600f548714158015610e8557506011548714155b8015610ea2575080516001600160581b0316610e9f613f69565b11155b8015610eb45750606081015160ff1615155b8015610ec257506080810151155b610f07576040805162461bcd60e51b81526020600482015260116024820152702bb4bd30b932103737ba103932b0b23c9760791b604482015290519081900360640190fd5b610f0f614f4a565b506000878152600d60209081526040918290208251808401909352805480845260019091015491830191909152610fcb5760408051808201825289815260208082018a815260008d8152600d83528490209251835551600190920191909155848201518483015183518d81529283018c905263ffffffff918216838501521660608201526080810189905290517f530a6602289d4bdb1c24e67332be34dd86bcf90f97911cf9ab69cf9d48db5eeb9181900360a00190a161108c565b805189141561103f5787891015610ff357610fed898989846020015188614036565b50611006565b611004888a83602001518a88614036565b505b6000888152600d6020526040812081815560010155831561103a576000898152601060205260408082208290558982528120555b61108c565b6040805162461bcd60e51b815260206004820181905260248201527f4f70706f6e656e742068617320612070656e64696e67206368616c6c656e6765604482015290519081900360640190fd5b505050505050505050565b60006001600160e01b031982166301ffc9a760e01b14806110c857506001600160e01b031982166317a0b21360e31b145b90505b919050565b6002546001600160a01b031681565b60085465ffffffffffff811691600160301b820463ffffffff90811692600160501b810490911691600160701b9091046001600160581b031690565b6000546001600160a01b031681565b6111326141c5565b80611170576040805162461bcd60e51b8152602060048201526009602482015268456d7074792069647360b81b604482015290519081900360640190fd5b60008060008484600081811061118257fe5b602090810292909201356000818152600b909352604090922054919250506001600160581b03166111b1613f69565b81106111f0576040805162461bcd60e51b81526020600482015260096024820152684e6f74206d6f6c647960b81b604482015290519081900360640190fd5b6000828152600b60205260409020600101541561123e576040805162461bcd60e51b81526020600482015260076024820152664475656c696e6760c81b604482015290519081900360640190fd5b60015b8581101561135b5786868281811061125557fe5b90506020020135945061126785613e02565b6000858152600b60205260409020546001600160581b03169350818410806112985750818414801561129857508285115b6112e9576040805162461bcd60e51b815260206004820152601c60248201527f57697a61726473206e6f74207374726963746c79206f72646572656400000000604482015290519081900360640190fd5b600581106112ff576112fa85614269565b61134d565b6000858152600b60205260409020600101541561134d576040805162461bcd60e51b81526020600482015260076024820152664475656c696e6760c81b604482015290519081900360640190fd5b849250839150600101611241565b50505050505050565b73673b537956a28e40aaa8d929fd1b6688c1583dda81565b6001546001600160a01b031633146113c6576040805162461bcd60e51b81526020600482015260086024820152674f6e6c7920434f4f60c01b604482015290519081900360640190fd5b600454600160601b900463ffffffff1680821115611424576040805162461bcd60e51b815260206004820152601660248201527524b73b30b634b2103830bab9b290323ab930ba34b7b760511b604482015290519081900360640190fd5b60035443838101918491600160301b90910465ffffffffffff1611156114b35760035465ffffffffffff808416600160301b909204161061149d576040805162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481c185d5cd95960921b604482015290519081900360640190fd5b50600354600160301b900465ffffffffffff1681035b6003805465ffffffffffff848116600160301b81026bffffffffffff000000000000198385168701841665ffffffffffff199586161781168217909555600480548085168801851690861617861682179055600580548085168801851690861617861682179055600680548085168801851690861617861682179055600780548085168801851690861617909516179093556008805480831686019092169190921617905560408051918252517f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e9181900360200190a150505050565b8161159a81613cc5565b816115a481613e02565b6115ac613c1f565b828414156115b957600080fd5b6115c28461194d565b80156115d257506115d28361194d565b6115db57600080fd5b6115e784846003614325565b50505050565b600854600090600160501b810463ffffffff90811660c802169065ffffffffffff16810143111561162257600091505061162a565b5050600c5415155b90565b6000546001600160a01b03163314611677576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b61168081614419565b600254604080516001600160a01b039283168152918316602083015280517fe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d69281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6116f2614486565b60008181526010602052604090205480611741576040805162461bcd60e51b815260206004820152600b60248201526a139bc81bdc1c1bdb995b9d60aa1b604482015290519081900360640190fd5b611749614f4a565b506000828152600d6020908152604091829020825180840190935280548084526001909101549183019190915282146117b5576040805162461bcd60e51b8152602060048201526009602482015268139bc818dbdb5b5a5d60ba1b604482015290519081900360640190fd5b6117c182846002614325565b506000918252600d602090815260408084208481556001018490556010909152808320839055908252812055565b60008060008060008060008060006118068a6136f2565b50975097509750975097509750975097506000600d60008c81526020019081526020016000206000015490508a898989898989898989604051602001808b81526020018a8152602001898152602001888152602001878152602001868152602001851515151560f81b8152600101848152602001831515151560f81b81526001018281526020019a5050505050505050505050604051602081830303815290604052805190602001209950505050505050505050919050565b6118c76141c5565b60005b818110156119485760008383838181106118e057fe5b602090810292909201356000818152600b90935260409092205491925050600160581b90046001600160581b03161580159061193157506000818152600b60205260409020546001600160581b0316155b1561193f5761193f81614269565b506001016118ca565b505050565b60008161195981613e02565b611961614f1c565b506000838152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff1660608301526001015460808201526119d1848261452a565b949350505050565b6119e161459c565b806119eb81613cc5565b6119f3614f1c565b506000828152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff166060830152600101546080820152611a63838261452a565b611ab4576040805162461bcd60e51b815260206004820152601b60248201527f43616e277420617363656e64206120627573792077697a617264210000000000604482015290519081900360640190fd5b611abc613f69565b60020281600001516001600160581b031610611b1f576040805162461bcd60e51b815260206004820152601a60248201527f4e6f7420656c696769626c6520666f7220617363656e73696f6e000000000000604482015290519081900360640190fd5b600f5415611b8c57600f805460009081526010602090815260408083208790559254868352918390208290558251918252810185905281517f317afd6701fe06472de1aa6af055861008f7e57c67729e0804c44be9a48facee929181900390910190a16000600f55611948565b600f8390556040805184815290517f75e589c241c6adb443794af5ecb220cdeb59eb1c44a37c6d1dfaebde22706b519181900360200190a1505050565b80611bd381613e02565b60007335b7838dd7507ada69610397a85310ae0abd50346001600160a01b031663fac8eafc846040518263ffffffff1660e01b81526004018082815260200191505060806040518083038186803b158015611c2d57600080fd5b505afa158015611c41573d6000803e3d6000fd5b505050506040513d6080811015611c5757600080fd5b506040908101516000858152600b60205291909120805491925090600160d01b900460ff1615611cce576040805162461bcd60e51b815260206004820152601860248201527f416666696e69747920616c726561647920757064617465640000000000000000604482015290519081900360640190fd5b805460ff909216600160d01b0260ff60d01b199092169190911790555050565b6000888152600b602090815260408083208a8452928190208354815483513060601b81870152603481018f9052605481018e90526001600160e01b0319600160b01b9384900460e090811b821660748401529390920490921b166078820152607c81018b9052609c8082018b90528351808303909101815260bc9091019092528151919092012060018301548114611dc1576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964206475656c206461746160781b604482015290519081900360640190fd5b6040805160208082018a90528183018890528251808303840181526060909201909252805191012089148015611e1d57506040805160208082018990528183018790528251808303840181526060909201909252805191012088145b611e6e576040805162461bcd60e51b815260206004820152601c60248201527f4d6f76657320646f6e2774206d6174636820636f6d6d69746d656e7400000000604482015290519081900360640190fd5b611e7b818c8c8a8a614640565b5050505050505050505050565b600a546001600160a01b031681565b611e9f614a41565b611ea76115ed565b15611eed576040805162461bcd60e51b8152602060048201526011602482015270546f75726e616d656e742061637469766560781b604482015290519081900360640190fd5b33ff5b80611efa81613cc5565b60008281526010602052604090205415611f5b576040805162461bcd60e51b815260206004820152601d60248201527f43616e27742063616e63656c20417363656e73696f6e20426174746c65000000604482015290519081900360640190fd5b6000828152600d602052604090205415611fa3576040805183815290517f04d33264c6183366c5711d870e69908ac063f2c36c434c4b1f77916c238b9dfc9181900360200190a15b506000908152600d6020526040812081815560010155565b6000546001600160a01b03163314612005576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b61200e81614419565b61201781614aa9565b50565b8061202481613e02565b61202c614b13565b61207d576040805162461bcd60e51b815260206004820152601a60248201527f4f6e6c7920647572696e67205265766976616c20506861736573000000000000604482015290519081900360640190fd5b612085614a41565b6000828152600b6020526040812080546009549192600160581b9091046001600160581b03169134816120b457fe5b0490506120bf613f69565b816001600160581b03161180156120e85750816001600160581b0316816001600160581b031611155b61212f576040805162461bcd60e51b8152602060048201526013602482015272125b9d985b1a59081c1bddd95c881b195d995b606a1b604482015290519081900360640190fd5b82546001600160581b03161561218c576040805162461bcd60e51b815260206004820152601d60248201527f43616e206f6e6c79207265766976652074697265642057697a61726473000000604482015290519081900360640190fd5b825463ffffffff600160b01b6001600160581b0384166001600160581b0319909316831781810483166001019092160263ffffffff60b01b1990911617845560408051878152602081019290925280517fa78677222d515efffcb323b960622c3e2bff3331916798f375b592c2a07f6c5a9281900390910190a15050505050565b60035460045460055460065460075465ffffffffffff80861696600160301b870482169663ffffffff600160601b8204811697600160801b808404831698600160a01b9094048316978683169792829004841696808616969583900485169581831695928490048316949182169390910490911690565b6000546001600160a01b031633146122ce576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b6122d781614419565b600154604080516001600160a01b039283168152918316602083015280517f1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe11849281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b612349613c1f565b600f5461235581613cc5565b601154806123a3576040805162461bcd60e51b8152602060048201526016602482015275139bc818da185b1b195b99d9481d1bc81858d8d95c1d60521b604482015290519081900360640190fd5b600f548110156123c8576123c281600f54601160010154866001614036565b506123df565b6123dd600f5482856011600101546001614036565b505b5050600060118190556012819055600f5550565b600c5490565b60095481565b612407614486565b600f54612451576040805162461bcd60e51b8152602060048201526013602482015272139bc815da5e985c99081d1bc8185cd8d95b99606a1b604482015290519081900360640190fd5b600f546000908152600b60205260409020601154156124825761247d600f546011600001546001614325565b6124c6565b805461249e90829060036001600160581b039182160216614bdc565b805463ffffffff600160b01b80830482166001019091160263ffffffff60b01b199091161781555b600f548154604080519283526001600160581b03909116602083015280517f5b0c8f5b5cdb91a5273ebf46ae7473ad46e9c17fb914ff3e5965ed6a0696c45f9281900390910190a1506000600f55565b6001546001600160a01b031681565b600354600160301b900465ffffffffffff16431090565b612544614c4e565b612595576040805162461bcd60e51b815260206004820152601860248201527f4f6e6c7920647572696e6720456e746572205068617365730000000000000000604482015290519081900360640190fd5b61259d614a41565b8281146125f1576040805162461bcd60e51b815260206004820152601c60248201527f4d69736d61746368656420706172616d65746572206c656e6774687300000000604482015290519081900360640190fd5b6000805b8481101561287b57600086868381811061260b57fe5b905060200201359050600085858481811061262257fe5b6000858152600b60209081526040909120546001600160581b0391909202939093013583169350600160581b90049091161590506126a7576040805162461bcd60e51b815260206004820152601c60248201527f57697a61726420616c726561647920696e20746f75726e616d656e7400000000604482015290519081900360640190fd5b6000807335b7838dd7507ada69610397a85310ae0abd50346001600160a01b031663fac8eafc856040518263ffffffff1660e01b81526004018082815260200191505060806040518083038186803b15801561270257600080fd5b505afa158015612716573d6000803e3d6000fd5b505050506040513d608081101561272c57600080fd5b50602081015160409091015190925090506001600160581b038316158015906127675750816001600160581b0316836001600160581b031611155b6127a8576040805162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103837bbb2b960991b604482015290519081900360640190fd5b6040805160a0810182526001600160581b039485168082526020808301828152600084860181815260ff97881660608701908152608087018381529b8352600b909452959020935184549151955192516001600160581b0319909216908916176affffffffffffffffffffff60581b1916600160581b95909816949094029690961763ffffffff60b01b1916600160b01b63ffffffff909716969096029590951760ff60d01b1916600160d01b9290931691909102919091178155925160019384015550600954029290920191016125f5565b50600c805485019055348111156128ce576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b5050505050565b600f5490565b60006128e5613c1f565b6128ee89613e02565b6128f788613e02565b87891061294b576040805162461bcd60e51b815260206004820152601a60248201527f57697a61726420494473206d757374206265206f726465726564000000000000604482015290519081900360640190fd5b600089815260106020526040812054151580612974575060008981526010602052604090205415155b15612a105760008a815260106020526040902054891480156129a357506000898152601060205260409020548a145b6129f4576040805162461bcd60e51b815260206004820152601d60248201527f4d757374207265736f6c766520417363656e73696f6e20426174746c65000000604482015290519081900360640190fd5b5060008981526010602052604080822082905589825281205560015b612a18614f1c565b5060008a8152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff166060830152600101546080820152612a86614f1c565b5060008a8152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff166060830152600101546080820152612af68c8361452a565b8015612b075750612b078b8261452a565b612b4b576040805162461bcd60e51b815260206004820152601060248201526f57697a617264206e6f7420726561647960801b604482015290519081900360640190fd5b6000612b628d8d856040015185604001518f614d07565b90506000612b7b8e8e866040015186604001518f614d07565b90507335b7838dd7507ada69610397a85310ae0abd50346001600160a01b031663a096d9f08f8f85858f8f8f8f6040518963ffffffff1660e01b81526004018089815260200188815260200187815260200186815260200180602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509a505050505050505050505060006040518083038186803b158015612c4d57600080fd5b505afa158015612c61573d6000803e3d6000fd5b50505050612c728e8e8e8e89614036565b9550600d60008f815260200190815260200160002060008082016000905560018201600090555050600d60008e815260200190815260200160002060008082016000905560018201600090555050505050505098975050505050505050565b7335b7838dd7507ada69610397a85310ae0abd503481565b80612cf381613e02565b612cfb6141c5565b6000612d05613f69565b6000848152600b60205260409020549091506001600160581b0316811115612d6b576040805162461bcd60e51b81526020600482015260146024820152735375727669766f722069736e277420616c69766560601b604482015290519081900360640190fd5b60005b84811015612ded576000868683818110612d8457fe5b602090810292909201356000818152600b90935260409092205491925050600160581b90046001600160581b031615801590612dd657506000818152600b60205260409020546001600160581b031683115b15612de457612de481614269565b50600101612d6e565b505050505050565b612dfd6141c5565b82612e0781613cc5565b6005600c541115612e53576040805162461bcd60e51b8152602060048201526011602482015270546f6f20736f6f6e20746f20636c61696d60781b604482015290519081900360640190fd5b600c548214612ea9576040805162461bcd60e51b815260206004820152601860248201527f4d7573742070726f7669646520616c6c2077696e6e6572730000000000000000604482015290519081900360640190fd5b6000848152600b60205260409020546001600160581b0316612f07576040805162461bcd60e51b81526020600482015260126024820152714e6f20636865657a6520666f7220796f752160701b604482015290519081900360640190fd5b6000612f11613f69565b9050600080805b8581101561306f576000878783818110612f2e57fe5b602090810292909201356000818152600b909352604090922054919250506001600160581b0316838211612fa9576040805162461bcd60e51b815260206004820152601e60248201527f57696e6e657273206e6f7420756e6971756520616e64206f7264657265640000604482015290519081900360640190fd5b6000828152600b6020526040902054600160581b90046001600160581b0316613019576040805162461bcd60e51b815260206004820152601960248201527f57697a61726420616c726561647920656c696d696e6174656400000000000000604482015290519081900360640190fd5b858110613060576040805162461bcd60e51b815260206004820152601060248201526f57697a617264206e6f74206d6f6c647960801b604482015290519081900360640190fd5b93909301929150600101612f18565b506000878152600b602052604081205483906001600160581b03163031028161309457fe5b60008a8152600b6020908152604080832080546001600160d81b031916815560010192909255600c805460001901905581518c81529390920491830182905280519193507fd53b67ba94a5d6268d11caa5d2693557779404ed02fc9825d86d2894d29cb8fd928290030190a1604051339082156108fc029083906000818181858888f1935050505015801561108c573d6000803e3d6000fd5b613135614f1c565b506000868152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff1660608301526001015460808201526131a3614f1c565b506000838152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff1660608301526001015460808083019190915282015180613257576040805162461bcd60e51b815260206004820152601260248201527157697a617264206e6f74206475656c696e6760701b604482015290519081900360640190fd5b6000858a10156132d557506040808401518382015182513060601b602080830191909152603482018e9052605482018a90526001600160e01b031960e094851b811660748401529290931b9091166078820152607c81018b9052609c8082018890528351808303909101815260bc9091019092528151910120613345565b506040808301518482015182513060601b602080830191909152603482018a9052605482018e90526001600160e01b031960e094851b811660748401529290931b9091166078820152607c8101879052609c8082018c90528351808303909101815260bc90910190925281519101205b81811461338d576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964206475656c206461746160781b604482015290519081900360640190fd5b6040805160208082018b90528183018a9052825180830384018152606090920190925280519101208914613408576040805162461bcd60e51b815260206004820152601c60248201527f4d6f76657320646f6e2774206d6174636820636f6d6d69746d656e7400000000604482015290519081900360640190fd5b600a5460408051630608feef60e21b8152600481018b905290516001600160a01b0390921691631823fbbc91602480820192602092909190829003018186803b15801561345457600080fd5b505afa158015613468573d6000803e3d6000fd5b505050506040513d602081101561347e57600080fd5b50516134c3576040805162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a59081b5bdd995cd95d608a1b604482015290519081900360640190fd5b6000828152600e602090815260408083208984529091529020541561354957858a101561351a576000828152600e602090815260408083208984529091529020546135159083908c9089908c90614640565b613544565b6000828152600e6020908152604080832089845290915290205461354490839088908d908c614640565b6135fe565b6000828152601360205260409020546001600160801b031643106135a3576040805162461bcd60e51b815260206004820152600c60248201526b111d595b08195e1c1a5c995960a21b604482015290519081900360640190fd5b6000828152600e602090815260408083208d84528252918290208a905581518481529081018c905280820188905290517f4761d35fc455d658f13ab27c916c9772625a2cd266c1e36b1e03b6bbe9e03f1a9181900360600190a15b50505050505050505050565b6136126141c5565b8061361c81613cc5565b600c54600114613664576040805162461bcd60e51b815260206004820152600e60248201526d4b656570206669676874696e672160901b604482015290519081900360640190fd5b604080518381523031602082015281517fd53b67ba94a5d6268d11caa5d2693557779404ed02fc9825d86d2894d29cb8fd929181900390910190a16000600c819055828152600b602052604080822080546001600160d81b0319168155600101829055513391303180156108fc02929091818181858888f19350505050158015611948573d6000803e3d6000fd5b60008060008060008060008060008961370a81613e02565b613712614f1c565b600b60008d81526020019081526020016000206040518060a00160405290816000820160009054906101000a90046001600160581b03166001600160581b03166001600160581b0316815260200160008201600b9054906101000a90046001600160581b03166001600160581b03166001600160581b031681526020016000820160169054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160008201601a9054906101000a900460ff1660ff1660ff1681526020016001820154815250509050806060015160ff169a5080600001516001600160581b0316995080602001516001600160581b03169850806040015163ffffffff169750806080015196508b600f54149550601060008d815260200190815260200160002054945080600001516001600160581b031661384e613f69565b11935061385b8c8261452a565b925050509193959799909294969850565b613874613c1f565b8161387e81613cc5565b601154156138d3576040805162461bcd60e51b815260206004820152601960248201527f57697a61726420616c7265616479206368616c6c656e67656400000000000000604482015290519081900360640190fd5b6138df83600f54613e6a565b6138e7614d7d565b61392d576040805162461bcd60e51b81526020600482015260126024820152714368616c6c656e676520746f6f206c61746560701b604482015290519081900360640190fd5b613935614f1c565b506000838152600b6020908152604091829020825160a08101845281546001600160581b038082168352600160581b82041693820193909352600160b01b830463ffffffff1693810193909352600160d01b90910460ff1660608301526001015460808201526139a5848261452a565b6139e9576040805162461bcd60e51b815260206004820152601060248201526f57697a617264206e6f7420726561647960801b604482015290519081900360640190fd5b604080518082018252858152602090810185905260118690556012859055600f54825190815290810186905280820185905290517f01b90f216a53259c7a1ed029ef4d7c085b483b5418be4845bcdc933dec1101039181900360600190a150505050565b6000828152600b6020526040808220838352912060018201548015801590613a785750808260010154145b613a8157600080fd5b6000818152601360205260409020546001600160801b0316431015613aa557600080fd5b815483546000838152600e602090815260408083208a84529091529020546001600160581b0391821692821692909201169015613af957613ae68482614bdc565b82546001600160581b0319168355613b31565b6000828152600e6020908152604080832088845290915290205415613b3157613b228382614bdc565b83546001600160581b03191684555b60006001808601829055848101829055855463ffffffff600160b01b808304821684018216810263ffffffff60b01b1993841617895587548181048316909401909116029116178455828152601360209081526040808320805470ffffffffffffffffffffffffffffffffff19169055600e82528083208984528083528184208490558884528252808320929092558554855483518681529283018a90528284018990526001600160581b03918216606084015216608082015290517fc7aa367585942c098842dd2573f26d69dabb55248bcaea3c4463cd96680e1c2e9181900360a00190a1505050505050565b6040805160808101825260055465ffffffffffff8082168352600160301b820416602083015263ffffffff600160601b8204811693830193909352600160801b90049091166060820152613c7290614e1e565b613cc3576040805162461bcd60e51b815260206004820152601860248201527f4f6e6c7920647572696e672046696768742057696e646f770000000000000000604482015290519081900360640190fd5b565b6000818152600b6020526040902054600160581b90046001600160581b0316613d2d576040805162461bcd60e51b815260206004820152601560248201527415da5e985c9908191bd95cc81b9bdd08195e1a5cdd605a1b604482015290519081900360640190fd5b6040805163430c208160e01b81523360048201526024810183905290517335b7838dd7507ada69610397a85310ae0abd50349163430c2081916044808301926020929190829003018186803b158015613d8557600080fd5b505afa158015613d99573d6000803e3d6000fd5b505050506040513d6020811015613daf57600080fd5b5051612017576040805162461bcd60e51b815260206004820152601960248201527f4d7573742062652057697a61726420636f6e74726f6c6c657200000000000000604482015290519081900360640190fd5b6000818152600b6020526040902054600160581b90046001600160581b0316612017576040805162461bcd60e51b815260206004820152601560248201527415da5e985c9908191bd95cc81b9bdd08195e1a5cdd605a1b604482015290519081900360640190fd5b6000828152600d602052604090205415613ecb576040805162461bcd60e51b815260206004820152601d60248201527f50656e64696e6720626174746c6520616c726561647920657869737473000000604482015290519081900360640190fd5b60008111613f19576040805162461bcd60e51b81526020600482015260166024820152754e6f2057697a61726420697320617363656e64696e6760501b604482015290519081900360640190fd5b80821415613f65576040805162461bcd60e51b815260206004820152601460248201527343616e6e6f74206475656c206f6e6573656c662160601b604482015290519081900360640190fd5b5050565b6000613f73614f61565b506040805160808101825260085465ffffffffffff811680835263ffffffff600160301b830481166020850152600160501b830416938301939093526001600160581b03600160701b909104166060820152904311613fe057606001516001600160581b0316905061162a565b6000816040015163ffffffff16826000015165ffffffffffff1643038161400357fe5b0490506058811115614013575060585b8082606001516001600160581b0316901b6001600160581b03169250505061162a565b6000858152600b602090815260408083208784528184208154815484513060601b81880152603481018d9052605481018c90526001600160e01b0319600160b01b9384900460e090811b821660748401529390920490921b166078820152607c8101899052609c8082018990528451808303909101815260bc90910190935282519290930191909120600180830182905583018190559290919084156140e5576140de614e8f565b90506140f9565b50600354600160a01b900463ffffffff1643015b6040805180820182526001600160801b038084168252871515602080840182815260008a8152601383528690209451855491511515600160801b0260ff60801b19919095166fffffffffffffffffffffffffffffffff1990921691909117169290921790925582518781529081018c90528083018b905260608101849052608081019190915260a0810189905260c0810188905290517f77d9948186effe9235713be25320cbda824d0fac31841e1af13e014cb0b348539181900360e00190a150505095945050505050565b6040805160808101825260075465ffffffffffff8082168352600160301b820416602083015263ffffffff600160601b8204811693830193909352600160801b9004909116606082015261421890614e1e565b613cc3576040805162461bcd60e51b815260206004820152601a60248201527f4f6e6c7920647572696e672043756c6c696e672057696e646f77000000000000604482015290519081900360640190fd5b6000818152600b6020526040902060010154156142c1576040805162461bcd60e51b815260206004820152601160248201527057697a617264206973206475656c696e6760781b604482015290519081900360640190fd5b6000818152600b6020908152604080832080546001600160d81b031916815560010192909255600c8054600019019055815183815291517f467c36a9931be143929e59fbcf25cc5ea17577e78af28da39a7b5f23d9c081b39281900390910190a150565b6000838152600b60205260408082208483529120815481546143589183916001600160581b039182169082160116614bdc565b815460408051878152602081018790526001600160581b039092168282015260ff85166060830152517f6f76f2ae8cea65409255207333bfb6f8c087e8e3b6b3f8cc7798fc35ed7af99a9181900360800190a1815479ffffffff0000000000000000000000ffffffffffffffffffffff1981166001600160b01b6001600160581b031990931683900463ffffffff90811682018116840292909217909455825463ffffffff60b01b1981169083900482169094011602919091179055505050565b6001600160a01b0381161580159061443f57506000546001600160a01b03828116911614155b612017576040805162461bcd60e51b8152602060048201526013602482015272496e76616c69642043454f206164647265737360681b604482015290519081900360640190fd5b6040805160808101825260065465ffffffffffff8082168352600160301b820416602083015263ffffffff600160601b8204811693830193909352600160801b900490911660608201526144d990614e1e565b613cc3576040805162461bcd60e51b815260206004820152601d60248201527f4f6e6c7920647572696e67205265736f6c7574696f6e2057696e646f77000000604482015290519081900360640190fd5b6000600f54831415801561454a5750600083815260106020526040902054155b801561455857506011548314155b8015614575575081516001600160581b0316614572613f69565b11155b80156145875750606082015160ff1615155b801561459557506080820151155b9392505050565b6040805160808101825260045465ffffffffffff8082168352600160301b820416602083015263ffffffff600160601b8204811693830193909352600160801b900490911660608201526145ef90614e1e565b613cc3576040805162461bcd60e51b815260206004820152601c60248201527f4f6e6c7920647572696e6720417363656e73696f6e2057696e646f7700000000604482015290519081900360640190fd5b614648614f4a565b506000858152601360209081526040918290208251808401909352546001600160801b038116808452600160801b90910460ff1615159183019190915243106146c7576040805162461bcd60e51b815260206004820152600c60248201526b111d595b08195e1c1a5c995960a21b604482015290519081900360640190fd5b6000858152600b602052604080822086835290822081548154929391926001600160581b03918216929116906146fb613f69565b905085602001511561472d5780600202820183131561471c5781925061472d565b80600202830182131561472d578291505b600a548554855460408051632c22625360e21b8152600481018d9052602481018c9052604481018890526064810187905260ff600160d01b94859004811660848301529390920490921660a482015290516000926001600160a01b03169163b089894c9160c4808301926020929190829003018186803b1580156147b057600080fd5b505afa1580156147c4573d6000803e3d6000fd5b505050506040513d60208110156147da57600080fd5b5051905060008490038112156147f557836000039050614800565b828113156148005750815b60208701519381019392819003921561482557828412614821578201614825565b8390035b855485546001600160581b03918216830191168290038382121561484c5760009101614858565b83811215614858570160005b6148628883614bdc565b61486c8782614bdc565b6000801b88600101819055506000801b876001018190555060018860000160168282829054906101000a900463ffffffff160192506101000a81548163ffffffff021916908363ffffffff16021790555060018760000160168282829054906101000a900463ffffffff160192506101000a81548163ffffffff021916908363ffffffff160217905550601360008f8152602001908152602001600020600080820160006101000a8154906001600160801b0302191690556000820160106101000a81549060ff02191690555050600e60008f815260200190815260200160002060008e815260200190815260200160002060009055600e60008f815260200190815260200160002060008d8152602001908152602001600020600090557f5d340f045b88c840424b9dc0322566ac25ac52a1710233de27fec1fcd04eae608e8e8e8e8e8d60000160009054906101000a90046001600160581b03168d60000160009054906101000a90046001600160581b031660405180888152602001878152602001868152602001858152602001848152602001836001600160581b03168152602001826001600160581b0316815260200197505050505050505060405180910390a15050505050505050505050505050565b3373673b537956a28e40aaa8d929fd1b6688c1583dda14613cc3576040805162461bcd60e51b815260206004820152601860248201527f4f6e6c7920476174654b65657065722063616e2063616c6c0000000000000000604482015290519081900360640190fd5b600054604080516001600160a01b039283168152918316602083015280517f9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc69281900390910190a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000614b1d614f1c565b506040805160a08101825260035465ffffffffffff8082168352600160301b8204166020830181905263ffffffff600160601b8304811694840194909452600160801b820484166060840152600160a01b909104909216608082015290431015614b8b57600091505061162a565b806040015163ffffffff1681600001510165ffffffffffff164310158015614bd65750806060015163ffffffff16816040015163ffffffff168260000151010165ffffffffffff1643105b91505090565b6001600160581b03811315614bf557506001600160581b035b81546001600160581b0319166001600160581b0382811691909117808455600160581b900416811315613f655781546001600160581b038216600160581b026affffffffffffffffffffff60581b199091161782555050565b6000614c58614f1c565b506040805160a08101825260035465ffffffffffff8082168352600160301b8204166020830181905263ffffffff600160601b8304811694840194909452600160801b820484166060840152600160a01b909104909216608082015290431015614cc657600091505061162a565b805165ffffffffffff164310801590614bd65750806060015163ffffffff16816040015163ffffffff168260000151010165ffffffffffff16431091505090565b60408051601960f81b602080830191909152600060218301523060601b602283015260368201889052605682018790526001600160e01b031960e087811b8216607685015286901b16607a830152607e80830185905283518084039091018152609e909201909252805191012095945050505050565b6000614d87614f61565b506040805160808101825260065465ffffffffffff808216808452600160301b8304909116602084015263ffffffff600160601b83048116948401859052600160801b909204909116606083015290916000914382010381614de557fe5b0490506000826040015163ffffffff168202836000015165ffffffffffff1601905043836060015163ffffffff16820311935050505090565b6000816020015165ffffffffffff16431015614e3c575060006110cb565b815165ffffffffffff16431015614e55575060006110cb565b6000826040015163ffffffff16836000015165ffffffffffff16430381614e7857fe5b606085015163ffffffff1691900610915050919050565b6000614e99614f61565b506040805160808101825260075465ffffffffffff808216808452600160301b8304909116602084015263ffffffff600160601b83048116948401859052600160801b909204909116606083015290916000914382010381614ef757fe5b049050816040015163ffffffff168102826000015165ffffffffffff16019250505090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080518082019091526000808252602082015290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a7231582033f208acfd29abed4471892c643e2e8fa01c65d0d0330bef064cae40a94ee97f64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059436f6e7374727563746f7220617267756d656e7473206d757374206265206e6f6e2d30"

// DeployBasicTournament deploys a new Ethereum contract, binding an instance of BasicTournament to it.
func DeployBasicTournament(auth *bind.TransactOpts, backend bind.ContractBackend, cooAddress_ common.Address, duelResolver_ common.Address, powerScale_ *big.Int, tournamentStartBlock_ *big.Int, admissionDuration_ uint32, revivalDuration_ uint32, ascensionDuration_ *big.Int, fightDuration_ *big.Int, cullingDuration_ *big.Int, blueMoldBasePower_ *big.Int, sessionsBetweenMoldDoubling_ *big.Int, duelTimeoutBlocks_ *big.Int) (common.Address, *types.Transaction, *BasicTournament, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTournamentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTournamentBin), backend, cooAddress_, duelResolver_, powerScale_, tournamentStartBlock_, admissionDuration_, revivalDuration_, ascensionDuration_, fightDuration_, cullingDuration_, blueMoldBasePower_, sessionsBetweenMoldDoubling_, duelTimeoutBlocks_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicTournament{BasicTournamentCaller: BasicTournamentCaller{contract: contract}, BasicTournamentTransactor: BasicTournamentTransactor{contract: contract}, BasicTournamentFilterer: BasicTournamentFilterer{contract: contract}}, nil
}

// BasicTournament is an auto generated Go binding around an Ethereum contract.
type BasicTournament struct {
	BasicTournamentCaller     // Read-only binding to the contract
	BasicTournamentTransactor // Write-only binding to the contract
	BasicTournamentFilterer   // Log filterer for contract events
}

// BasicTournamentCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTournamentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTournamentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTournamentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTournamentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTournamentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTournamentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTournamentSession struct {
	Contract     *BasicTournament  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTournamentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTournamentCallerSession struct {
	Contract *BasicTournamentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BasicTournamentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTournamentTransactorSession struct {
	Contract     *BasicTournamentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BasicTournamentRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTournamentRaw struct {
	Contract *BasicTournament // Generic contract binding to access the raw methods on
}

// BasicTournamentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTournamentCallerRaw struct {
	Contract *BasicTournamentCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTournamentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTournamentTransactorRaw struct {
	Contract *BasicTournamentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicTournament creates a new instance of BasicTournament, bound to a specific deployed contract.
func NewBasicTournament(address common.Address, backend bind.ContractBackend) (*BasicTournament, error) {
	contract, err := bindBasicTournament(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicTournament{BasicTournamentCaller: BasicTournamentCaller{contract: contract}, BasicTournamentTransactor: BasicTournamentTransactor{contract: contract}, BasicTournamentFilterer: BasicTournamentFilterer{contract: contract}}, nil
}

// NewBasicTournamentCaller creates a new read-only instance of BasicTournament, bound to a specific deployed contract.
func NewBasicTournamentCaller(address common.Address, caller bind.ContractCaller) (*BasicTournamentCaller, error) {
	contract, err := bindBasicTournament(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTournamentCaller{contract: contract}, nil
}

// NewBasicTournamentTransactor creates a new write-only instance of BasicTournament, bound to a specific deployed contract.
func NewBasicTournamentTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTournamentTransactor, error) {
	contract, err := bindBasicTournament(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTournamentTransactor{contract: contract}, nil
}

// NewBasicTournamentFilterer creates a new log filterer instance of BasicTournament, bound to a specific deployed contract.
func NewBasicTournamentFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTournamentFilterer, error) {
	contract, err := bindBasicTournament(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTournamentFilterer{contract: contract}, nil
}

// bindBasicTournament binds a generic wrapper to an already deployed contract.
func bindBasicTournament(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTournamentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicTournament *BasicTournamentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicTournament.Contract.BasicTournamentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicTournament *BasicTournamentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTournament.Contract.BasicTournamentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicTournament *BasicTournamentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicTournament.Contract.BasicTournamentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicTournament *BasicTournamentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicTournament.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicTournament *BasicTournamentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTournament.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicTournament *BasicTournamentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicTournament.Contract.contract.Transact(opts, method, params...)
}

// GATEKEEPER is a free data retrieval call binding the contract method 0x128183ba.
//
// Solidity: function GATE_KEEPER() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) GATEKEEPER(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "GATE_KEEPER")
	return *ret0, err
}

// GATEKEEPER is a free data retrieval call binding the contract method 0x128183ba.
//
// Solidity: function GATE_KEEPER() constant returns(address)
func (_BasicTournament *BasicTournamentSession) GATEKEEPER() (common.Address, error) {
	return _BasicTournament.Contract.GATEKEEPER(&_BasicTournament.CallOpts)
}

// GATEKEEPER is a free data retrieval call binding the contract method 0x128183ba.
//
// Solidity: function GATE_KEEPER() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) GATEKEEPER() (common.Address, error) {
	return _BasicTournament.Contract.GATEKEEPER(&_BasicTournament.CallOpts)
}

// WIZARDGUILD is a free data retrieval call binding the contract method 0xc9eb068b.
//
// Solidity: function WIZARD_GUILD() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) WIZARDGUILD(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "WIZARD_GUILD")
	return *ret0, err
}

// WIZARDGUILD is a free data retrieval call binding the contract method 0xc9eb068b.
//
// Solidity: function WIZARD_GUILD() constant returns(address)
func (_BasicTournament *BasicTournamentSession) WIZARDGUILD() (common.Address, error) {
	return _BasicTournament.Contract.WIZARDGUILD(&_BasicTournament.CallOpts)
}

// WIZARDGUILD is a free data retrieval call binding the contract method 0xc9eb068b.
//
// Solidity: function WIZARD_GUILD() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) WIZARDGUILD() (common.Address, error) {
	return _BasicTournament.Contract.WIZARDGUILD(&_BasicTournament.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentSession) CeoAddress() (common.Address, error) {
	return _BasicTournament.Contract.CeoAddress(&_BasicTournament.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) CeoAddress() (common.Address, error) {
	return _BasicTournament.Contract.CeoAddress(&_BasicTournament.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentSession) CfoAddress() (common.Address, error) {
	return _BasicTournament.Contract.CfoAddress(&_BasicTournament.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) CfoAddress() (common.Address, error) {
	return _BasicTournament.Contract.CfoAddress(&_BasicTournament.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_BasicTournament *BasicTournamentSession) CooAddress() (common.Address, error) {
	return _BasicTournament.Contract.CooAddress(&_BasicTournament.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) CooAddress() (common.Address, error) {
	return _BasicTournament.Contract.CooAddress(&_BasicTournament.CallOpts)
}

// DuelResolver is a free data retrieval call binding the contract method 0x810306b9.
//
// Solidity: function duelResolver() constant returns(address)
func (_BasicTournament *BasicTournamentCaller) DuelResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "duelResolver")
	return *ret0, err
}

// DuelResolver is a free data retrieval call binding the contract method 0x810306b9.
//
// Solidity: function duelResolver() constant returns(address)
func (_BasicTournament *BasicTournamentSession) DuelResolver() (common.Address, error) {
	return _BasicTournament.Contract.DuelResolver(&_BasicTournament.CallOpts)
}

// DuelResolver is a free data retrieval call binding the contract method 0x810306b9.
//
// Solidity: function duelResolver() constant returns(address)
func (_BasicTournament *BasicTournamentCallerSession) DuelResolver() (common.Address, error) {
	return _BasicTournament.Contract.DuelResolver(&_BasicTournament.CallOpts)
}

// GetAscendingWizardId is a free data retrieval call binding the contract method 0xbfd31c19.
//
// Solidity: function getAscendingWizardId() constant returns(uint256)
func (_BasicTournament *BasicTournamentCaller) GetAscendingWizardId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "getAscendingWizardId")
	return *ret0, err
}

// GetAscendingWizardId is a free data retrieval call binding the contract method 0xbfd31c19.
//
// Solidity: function getAscendingWizardId() constant returns(uint256)
func (_BasicTournament *BasicTournamentSession) GetAscendingWizardId() (*big.Int, error) {
	return _BasicTournament.Contract.GetAscendingWizardId(&_BasicTournament.CallOpts)
}

// GetAscendingWizardId is a free data retrieval call binding the contract method 0xbfd31c19.
//
// Solidity: function getAscendingWizardId() constant returns(uint256)
func (_BasicTournament *BasicTournamentCallerSession) GetAscendingWizardId() (*big.Int, error) {
	return _BasicTournament.Contract.GetAscendingWizardId(&_BasicTournament.CallOpts)
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_BasicTournament *BasicTournamentCaller) GetBlueMoldParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _BasicTournament.contract.Call(opts, out, "getBlueMoldParameters")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_BasicTournament *BasicTournamentSession) GetBlueMoldParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BasicTournament.Contract.GetBlueMoldParameters(&_BasicTournament.CallOpts)
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_BasicTournament *BasicTournamentCallerSession) GetBlueMoldParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BasicTournament.Contract.GetBlueMoldParameters(&_BasicTournament.CallOpts)
}

// GetRemainingWizards is a free data retrieval call binding the contract method 0x9ce0c954.
//
// Solidity: function getRemainingWizards() constant returns(uint256)
func (_BasicTournament *BasicTournamentCaller) GetRemainingWizards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "getRemainingWizards")
	return *ret0, err
}

// GetRemainingWizards is a free data retrieval call binding the contract method 0x9ce0c954.
//
// Solidity: function getRemainingWizards() constant returns(uint256)
func (_BasicTournament *BasicTournamentSession) GetRemainingWizards() (*big.Int, error) {
	return _BasicTournament.Contract.GetRemainingWizards(&_BasicTournament.CallOpts)
}

// GetRemainingWizards is a free data retrieval call binding the contract method 0x9ce0c954.
//
// Solidity: function getRemainingWizards() constant returns(uint256)
func (_BasicTournament *BasicTournamentCallerSession) GetRemainingWizards() (*big.Int, error) {
	return _BasicTournament.Contract.GetRemainingWizards(&_BasicTournament.CallOpts)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_BasicTournament *BasicTournamentCaller) GetTimeParameters(opts *bind.CallOpts) (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	ret := new(struct {
		TournamentStartBlock     *big.Int
		PauseEndedBlock          *big.Int
		AdmissionDuration        *big.Int
		RevivalDuration          *big.Int
		DuelTimeoutDuration      *big.Int
		AscensionWindowStart     *big.Int
		AscensionWindowDuration  *big.Int
		FightWindowStart         *big.Int
		FightWindowDuration      *big.Int
		ResolutionWindowStart    *big.Int
		ResolutionWindowDuration *big.Int
		CullingWindowStart       *big.Int
		CullingWindowDuration    *big.Int
	})
	out := ret
	err := _BasicTournament.contract.Call(opts, out, "getTimeParameters")
	return *ret, err
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_BasicTournament *BasicTournamentSession) GetTimeParameters() (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	return _BasicTournament.Contract.GetTimeParameters(&_BasicTournament.CallOpts)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_BasicTournament *BasicTournamentCallerSession) GetTimeParameters() (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	return _BasicTournament.Contract.GetTimeParameters(&_BasicTournament.CallOpts)
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 wizardId) constant returns(uint256 affinity, uint256 power, uint256 maxPower, uint256 nonce, bytes32 currentDuel, bool ascending, uint256 ascensionOpponent, bool molded, bool ready)
func (_BasicTournament *BasicTournamentCaller) GetWizard(opts *bind.CallOpts, wizardId *big.Int) (struct {
	Affinity          *big.Int
	Power             *big.Int
	MaxPower          *big.Int
	Nonce             *big.Int
	CurrentDuel       [32]byte
	Ascending         bool
	AscensionOpponent *big.Int
	Molded            bool
	Ready             bool
}, error) {
	ret := new(struct {
		Affinity          *big.Int
		Power             *big.Int
		MaxPower          *big.Int
		Nonce             *big.Int
		CurrentDuel       [32]byte
		Ascending         bool
		AscensionOpponent *big.Int
		Molded            bool
		Ready             bool
	})
	out := ret
	err := _BasicTournament.contract.Call(opts, out, "getWizard", wizardId)
	return *ret, err
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 wizardId) constant returns(uint256 affinity, uint256 power, uint256 maxPower, uint256 nonce, bytes32 currentDuel, bool ascending, uint256 ascensionOpponent, bool molded, bool ready)
func (_BasicTournament *BasicTournamentSession) GetWizard(wizardId *big.Int) (struct {
	Affinity          *big.Int
	Power             *big.Int
	MaxPower          *big.Int
	Nonce             *big.Int
	CurrentDuel       [32]byte
	Ascending         bool
	AscensionOpponent *big.Int
	Molded            bool
	Ready             bool
}, error) {
	return _BasicTournament.Contract.GetWizard(&_BasicTournament.CallOpts, wizardId)
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 wizardId) constant returns(uint256 affinity, uint256 power, uint256 maxPower, uint256 nonce, bytes32 currentDuel, bool ascending, uint256 ascensionOpponent, bool molded, bool ready)
func (_BasicTournament *BasicTournamentCallerSession) GetWizard(wizardId *big.Int) (struct {
	Affinity          *big.Int
	Power             *big.Int
	MaxPower          *big.Int
	Nonce             *big.Int
	CurrentDuel       [32]byte
	Ascending         bool
	AscensionOpponent *big.Int
	Molded            bool
	Ready             bool
}, error) {
	return _BasicTournament.Contract.GetWizard(&_BasicTournament.CallOpts, wizardId)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_BasicTournament *BasicTournamentCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "isActive")
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_BasicTournament *BasicTournamentSession) IsActive() (bool, error) {
	return _BasicTournament.Contract.IsActive(&_BasicTournament.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_BasicTournament *BasicTournamentCallerSession) IsActive() (bool, error) {
	return _BasicTournament.Contract.IsActive(&_BasicTournament.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_BasicTournament *BasicTournamentCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "isPaused")
	return *ret0, err
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_BasicTournament *BasicTournamentSession) IsPaused() (bool, error) {
	return _BasicTournament.Contract.IsPaused(&_BasicTournament.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_BasicTournament *BasicTournamentCallerSession) IsPaused() (bool, error) {
	return _BasicTournament.Contract.IsPaused(&_BasicTournament.CallOpts)
}

// IsReady is a free data retrieval call binding the contract method 0x50df8f71.
//
// Solidity: function isReady(uint256 wizardId) constant returns(bool)
func (_BasicTournament *BasicTournamentCaller) IsReady(opts *bind.CallOpts, wizardId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "isReady", wizardId)
	return *ret0, err
}

// IsReady is a free data retrieval call binding the contract method 0x50df8f71.
//
// Solidity: function isReady(uint256 wizardId) constant returns(bool)
func (_BasicTournament *BasicTournamentSession) IsReady(wizardId *big.Int) (bool, error) {
	return _BasicTournament.Contract.IsReady(&_BasicTournament.CallOpts, wizardId)
}

// IsReady is a free data retrieval call binding the contract method 0x50df8f71.
//
// Solidity: function isReady(uint256 wizardId) constant returns(bool)
func (_BasicTournament *BasicTournamentCallerSession) IsReady(wizardId *big.Int) (bool, error) {
	return _BasicTournament.Contract.IsReady(&_BasicTournament.CallOpts, wizardId)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_BasicTournament *BasicTournamentCaller) PowerScale(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "powerScale")
	return *ret0, err
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_BasicTournament *BasicTournamentSession) PowerScale() (*big.Int, error) {
	return _BasicTournament.Contract.PowerScale(&_BasicTournament.CallOpts)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_BasicTournament *BasicTournamentCallerSession) PowerScale() (*big.Int, error) {
	return _BasicTournament.Contract.PowerScale(&_BasicTournament.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_BasicTournament *BasicTournamentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_BasicTournament *BasicTournamentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasicTournament.Contract.SupportsInterface(&_BasicTournament.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_BasicTournament *BasicTournamentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasicTournament.Contract.SupportsInterface(&_BasicTournament.CallOpts, interfaceId)
}

// WizardFingerprint is a free data retrieval call binding the contract method 0x3f976ca9.
//
// Solidity: function wizardFingerprint(uint256 wizardId) constant returns(bytes32)
func (_BasicTournament *BasicTournamentCaller) WizardFingerprint(opts *bind.CallOpts, wizardId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasicTournament.contract.Call(opts, out, "wizardFingerprint", wizardId)
	return *ret0, err
}

// WizardFingerprint is a free data retrieval call binding the contract method 0x3f976ca9.
//
// Solidity: function wizardFingerprint(uint256 wizardId) constant returns(bytes32)
func (_BasicTournament *BasicTournamentSession) WizardFingerprint(wizardId *big.Int) ([32]byte, error) {
	return _BasicTournament.Contract.WizardFingerprint(&_BasicTournament.CallOpts, wizardId)
}

// WizardFingerprint is a free data retrieval call binding the contract method 0x3f976ca9.
//
// Solidity: function wizardFingerprint(uint256 wizardId) constant returns(bytes32)
func (_BasicTournament *BasicTournamentCallerSession) WizardFingerprint(wizardId *big.Int) ([32]byte, error) {
	return _BasicTournament.Contract.WizardFingerprint(&_BasicTournament.CallOpts, wizardId)
}

// AcceptAscensionChallenge is a paid mutator transaction binding the contract method 0x9a963fcd.
//
// Solidity: function acceptAscensionChallenge(bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactor) AcceptAscensionChallenge(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "acceptAscensionChallenge", commitment)
}

// AcceptAscensionChallenge is a paid mutator transaction binding the contract method 0x9a963fcd.
//
// Solidity: function acceptAscensionChallenge(bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentSession) AcceptAscensionChallenge(commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.AcceptAscensionChallenge(&_BasicTournament.TransactOpts, commitment)
}

// AcceptAscensionChallenge is a paid mutator transaction binding the contract method 0x9a963fcd.
//
// Solidity: function acceptAscensionChallenge(bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactorSession) AcceptAscensionChallenge(commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.AcceptAscensionChallenge(&_BasicTournament.TransactOpts, commitment)
}

// CancelCommitment is a paid mutator transaction binding the contract method 0x840a1ff4.
//
// Solidity: function cancelCommitment(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) CancelCommitment(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "cancelCommitment", wizardId)
}

// CancelCommitment is a paid mutator transaction binding the contract method 0x840a1ff4.
//
// Solidity: function cancelCommitment(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentSession) CancelCommitment(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CancelCommitment(&_BasicTournament.TransactOpts, wizardId)
}

// CancelCommitment is a paid mutator transaction binding the contract method 0x840a1ff4.
//
// Solidity: function cancelCommitment(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) CancelCommitment(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CancelCommitment(&_BasicTournament.TransactOpts, wizardId)
}

// ChallengeAscending is a paid mutator transaction binding the contract method 0xfb3790c5.
//
// Solidity: function challengeAscending(uint256 wizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactor) ChallengeAscending(opts *bind.TransactOpts, wizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "challengeAscending", wizardId, commitment)
}

// ChallengeAscending is a paid mutator transaction binding the contract method 0xfb3790c5.
//
// Solidity: function challengeAscending(uint256 wizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentSession) ChallengeAscending(wizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.ChallengeAscending(&_BasicTournament.TransactOpts, wizardId, commitment)
}

// ChallengeAscending is a paid mutator transaction binding the contract method 0xfb3790c5.
//
// Solidity: function challengeAscending(uint256 wizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactorSession) ChallengeAscending(wizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.ChallengeAscending(&_BasicTournament.TransactOpts, wizardId, commitment)
}

// ClaimSharedWinnings is a paid mutator transaction binding the contract method 0xcc4db960.
//
// Solidity: function claimSharedWinnings(uint256 claimingWinnerId, uint256[] allWinners) returns()
func (_BasicTournament *BasicTournamentTransactor) ClaimSharedWinnings(opts *bind.TransactOpts, claimingWinnerId *big.Int, allWinners []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "claimSharedWinnings", claimingWinnerId, allWinners)
}

// ClaimSharedWinnings is a paid mutator transaction binding the contract method 0xcc4db960.
//
// Solidity: function claimSharedWinnings(uint256 claimingWinnerId, uint256[] allWinners) returns()
func (_BasicTournament *BasicTournamentSession) ClaimSharedWinnings(claimingWinnerId *big.Int, allWinners []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ClaimSharedWinnings(&_BasicTournament.TransactOpts, claimingWinnerId, allWinners)
}

// ClaimSharedWinnings is a paid mutator transaction binding the contract method 0xcc4db960.
//
// Solidity: function claimSharedWinnings(uint256 claimingWinnerId, uint256[] allWinners) returns()
func (_BasicTournament *BasicTournamentTransactorSession) ClaimSharedWinnings(claimingWinnerId *big.Int, allWinners []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ClaimSharedWinnings(&_BasicTournament.TransactOpts, claimingWinnerId, allWinners)
}

// ClaimTheBigCheeze is a paid mutator transaction binding the contract method 0xda0cb2ae.
//
// Solidity: function claimTheBigCheeze(uint256 claimingWinnerId) returns()
func (_BasicTournament *BasicTournamentTransactor) ClaimTheBigCheeze(opts *bind.TransactOpts, claimingWinnerId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "claimTheBigCheeze", claimingWinnerId)
}

// ClaimTheBigCheeze is a paid mutator transaction binding the contract method 0xda0cb2ae.
//
// Solidity: function claimTheBigCheeze(uint256 claimingWinnerId) returns()
func (_BasicTournament *BasicTournamentSession) ClaimTheBigCheeze(claimingWinnerId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ClaimTheBigCheeze(&_BasicTournament.TransactOpts, claimingWinnerId)
}

// ClaimTheBigCheeze is a paid mutator transaction binding the contract method 0xda0cb2ae.
//
// Solidity: function claimTheBigCheeze(uint256 claimingWinnerId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) ClaimTheBigCheeze(claimingWinnerId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ClaimTheBigCheeze(&_BasicTournament.TransactOpts, claimingWinnerId)
}

// CompleteAscension is a paid mutator transaction binding the contract method 0xaedb27fc.
//
// Solidity: function completeAscension() returns()
func (_BasicTournament *BasicTournamentTransactor) CompleteAscension(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "completeAscension")
}

// CompleteAscension is a paid mutator transaction binding the contract method 0xaedb27fc.
//
// Solidity: function completeAscension() returns()
func (_BasicTournament *BasicTournamentSession) CompleteAscension() (*types.Transaction, error) {
	return _BasicTournament.Contract.CompleteAscension(&_BasicTournament.TransactOpts)
}

// CompleteAscension is a paid mutator transaction binding the contract method 0xaedb27fc.
//
// Solidity: function completeAscension() returns()
func (_BasicTournament *BasicTournamentTransactorSession) CompleteAscension() (*types.Transaction, error) {
	return _BasicTournament.Contract.CompleteAscension(&_BasicTournament.TransactOpts)
}

// CullMoldedWithMolded is a paid mutator transaction binding the contract method 0x0af29b96.
//
// Solidity: function cullMoldedWithMolded(uint256[] moldyWizardIds) returns()
func (_BasicTournament *BasicTournamentTransactor) CullMoldedWithMolded(opts *bind.TransactOpts, moldyWizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "cullMoldedWithMolded", moldyWizardIds)
}

// CullMoldedWithMolded is a paid mutator transaction binding the contract method 0x0af29b96.
//
// Solidity: function cullMoldedWithMolded(uint256[] moldyWizardIds) returns()
func (_BasicTournament *BasicTournamentSession) CullMoldedWithMolded(moldyWizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullMoldedWithMolded(&_BasicTournament.TransactOpts, moldyWizardIds)
}

// CullMoldedWithMolded is a paid mutator transaction binding the contract method 0x0af29b96.
//
// Solidity: function cullMoldedWithMolded(uint256[] moldyWizardIds) returns()
func (_BasicTournament *BasicTournamentTransactorSession) CullMoldedWithMolded(moldyWizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullMoldedWithMolded(&_BasicTournament.TransactOpts, moldyWizardIds)
}

// CullMoldedWithSurvivor is a paid mutator transaction binding the contract method 0xcbe6549e.
//
// Solidity: function cullMoldedWithSurvivor(uint256[] wizardIds, uint256 survivor) returns()
func (_BasicTournament *BasicTournamentTransactor) CullMoldedWithSurvivor(opts *bind.TransactOpts, wizardIds []*big.Int, survivor *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "cullMoldedWithSurvivor", wizardIds, survivor)
}

// CullMoldedWithSurvivor is a paid mutator transaction binding the contract method 0xcbe6549e.
//
// Solidity: function cullMoldedWithSurvivor(uint256[] wizardIds, uint256 survivor) returns()
func (_BasicTournament *BasicTournamentSession) CullMoldedWithSurvivor(wizardIds []*big.Int, survivor *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullMoldedWithSurvivor(&_BasicTournament.TransactOpts, wizardIds, survivor)
}

// CullMoldedWithSurvivor is a paid mutator transaction binding the contract method 0xcbe6549e.
//
// Solidity: function cullMoldedWithSurvivor(uint256[] wizardIds, uint256 survivor) returns()
func (_BasicTournament *BasicTournamentTransactorSession) CullMoldedWithSurvivor(wizardIds []*big.Int, survivor *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullMoldedWithSurvivor(&_BasicTournament.TransactOpts, wizardIds, survivor)
}

// CullTiredWizards is a paid mutator transaction binding the contract method 0x43d9922f.
//
// Solidity: function cullTiredWizards(uint256[] wizardIds) returns()
func (_BasicTournament *BasicTournamentTransactor) CullTiredWizards(opts *bind.TransactOpts, wizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "cullTiredWizards", wizardIds)
}

// CullTiredWizards is a paid mutator transaction binding the contract method 0x43d9922f.
//
// Solidity: function cullTiredWizards(uint256[] wizardIds) returns()
func (_BasicTournament *BasicTournamentSession) CullTiredWizards(wizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullTiredWizards(&_BasicTournament.TransactOpts, wizardIds)
}

// CullTiredWizards is a paid mutator transaction binding the contract method 0x43d9922f.
//
// Solidity: function cullTiredWizards(uint256[] wizardIds) returns()
func (_BasicTournament *BasicTournamentTransactorSession) CullTiredWizards(wizardIds []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.CullTiredWizards(&_BasicTournament.TransactOpts, wizardIds)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BasicTournament *BasicTournamentTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BasicTournament *BasicTournamentSession) Destroy() (*types.Transaction, error) {
	return _BasicTournament.Contract.Destroy(&_BasicTournament.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_BasicTournament *BasicTournamentTransactorSession) Destroy() (*types.Transaction, error) {
	return _BasicTournament.Contract.Destroy(&_BasicTournament.TransactOpts)
}

// DoubleCommit is a paid mutator transaction binding the contract method 0xc51b58aa.
//
// Solidity: function doubleCommit(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes sig1, bytes sig2) returns(bytes32 duelId)
func (_BasicTournament *BasicTournamentTransactor) DoubleCommit(opts *bind.TransactOpts, wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, sig1 []byte, sig2 []byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "doubleCommit", wizardId1, wizardId2, commit1, commit2, sig1, sig2)
}

// DoubleCommit is a paid mutator transaction binding the contract method 0xc51b58aa.
//
// Solidity: function doubleCommit(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes sig1, bytes sig2) returns(bytes32 duelId)
func (_BasicTournament *BasicTournamentSession) DoubleCommit(wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, sig1 []byte, sig2 []byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.DoubleCommit(&_BasicTournament.TransactOpts, wizardId1, wizardId2, commit1, commit2, sig1, sig2)
}

// DoubleCommit is a paid mutator transaction binding the contract method 0xc51b58aa.
//
// Solidity: function doubleCommit(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes sig1, bytes sig2) returns(bytes32 duelId)
func (_BasicTournament *BasicTournamentTransactorSession) DoubleCommit(wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, sig1 []byte, sig2 []byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.DoubleCommit(&_BasicTournament.TransactOpts, wizardId1, wizardId2, commit1, commit2, sig1, sig2)
}

// DoubleReveal is a paid mutator transaction binding the contract method 0x7aa66a7c.
//
// Solidity: function doubleReveal(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes32 moveSet1, bytes32 moveSet2, bytes32 salt1, bytes32 salt2) returns()
func (_BasicTournament *BasicTournamentTransactor) DoubleReveal(opts *bind.TransactOpts, wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, moveSet1 [32]byte, moveSet2 [32]byte, salt1 [32]byte, salt2 [32]byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "doubleReveal", wizardId1, wizardId2, commit1, commit2, moveSet1, moveSet2, salt1, salt2)
}

// DoubleReveal is a paid mutator transaction binding the contract method 0x7aa66a7c.
//
// Solidity: function doubleReveal(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes32 moveSet1, bytes32 moveSet2, bytes32 salt1, bytes32 salt2) returns()
func (_BasicTournament *BasicTournamentSession) DoubleReveal(wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, moveSet1 [32]byte, moveSet2 [32]byte, salt1 [32]byte, salt2 [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.DoubleReveal(&_BasicTournament.TransactOpts, wizardId1, wizardId2, commit1, commit2, moveSet1, moveSet2, salt1, salt2)
}

// DoubleReveal is a paid mutator transaction binding the contract method 0x7aa66a7c.
//
// Solidity: function doubleReveal(uint256 wizardId1, uint256 wizardId2, bytes32 commit1, bytes32 commit2, bytes32 moveSet1, bytes32 moveSet2, bytes32 salt1, bytes32 salt2) returns()
func (_BasicTournament *BasicTournamentTransactorSession) DoubleReveal(wizardId1 *big.Int, wizardId2 *big.Int, commit1 [32]byte, commit2 [32]byte, moveSet1 [32]byte, moveSet2 [32]byte, salt1 [32]byte, salt2 [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.DoubleReveal(&_BasicTournament.TransactOpts, wizardId1, wizardId2, commit1, commit2, moveSet1, moveSet2, salt1, salt2)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_BasicTournament *BasicTournamentTransactor) EnterWizards(opts *bind.TransactOpts, wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "enterWizards", wizardIds, powers)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_BasicTournament *BasicTournamentSession) EnterWizards(wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.EnterWizards(&_BasicTournament.TransactOpts, wizardIds, powers)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_BasicTournament *BasicTournamentTransactorSession) EnterWizards(wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.EnterWizards(&_BasicTournament.TransactOpts, wizardIds, powers)
}

// GiftPower is a paid mutator transaction binding the contract method 0x170f8cc8.
//
// Solidity: function giftPower(uint256 sendingWizardId, uint256 receivingWizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) GiftPower(opts *bind.TransactOpts, sendingWizardId *big.Int, receivingWizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "giftPower", sendingWizardId, receivingWizardId)
}

// GiftPower is a paid mutator transaction binding the contract method 0x170f8cc8.
//
// Solidity: function giftPower(uint256 sendingWizardId, uint256 receivingWizardId) returns()
func (_BasicTournament *BasicTournamentSession) GiftPower(sendingWizardId *big.Int, receivingWizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.GiftPower(&_BasicTournament.TransactOpts, sendingWizardId, receivingWizardId)
}

// GiftPower is a paid mutator transaction binding the contract method 0x170f8cc8.
//
// Solidity: function giftPower(uint256 sendingWizardId, uint256 receivingWizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) GiftPower(sendingWizardId *big.Int, receivingWizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.GiftPower(&_BasicTournament.TransactOpts, sendingWizardId, receivingWizardId)
}

// OneSidedCommit is a paid mutator transaction binding the contract method 0x01d1c810.
//
// Solidity: function oneSidedCommit(uint256 committingWizardId, uint256 otherWizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactor) OneSidedCommit(opts *bind.TransactOpts, committingWizardId *big.Int, otherWizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "oneSidedCommit", committingWizardId, otherWizardId, commitment)
}

// OneSidedCommit is a paid mutator transaction binding the contract method 0x01d1c810.
//
// Solidity: function oneSidedCommit(uint256 committingWizardId, uint256 otherWizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentSession) OneSidedCommit(committingWizardId *big.Int, otherWizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.OneSidedCommit(&_BasicTournament.TransactOpts, committingWizardId, otherWizardId, commitment)
}

// OneSidedCommit is a paid mutator transaction binding the contract method 0x01d1c810.
//
// Solidity: function oneSidedCommit(uint256 committingWizardId, uint256 otherWizardId, bytes32 commitment) returns()
func (_BasicTournament *BasicTournamentTransactorSession) OneSidedCommit(committingWizardId *big.Int, otherWizardId *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.OneSidedCommit(&_BasicTournament.TransactOpts, committingWizardId, otherWizardId, commitment)
}

// OneSidedReveal is a paid mutator transaction binding the contract method 0xcef9b488.
//
// Solidity: function oneSidedReveal(uint256 committingWizardId, bytes32 commit, bytes32 moveSet, bytes32 salt, uint256 otherWizardId, bytes32 otherCommit) returns()
func (_BasicTournament *BasicTournamentTransactor) OneSidedReveal(opts *bind.TransactOpts, committingWizardId *big.Int, commit [32]byte, moveSet [32]byte, salt [32]byte, otherWizardId *big.Int, otherCommit [32]byte) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "oneSidedReveal", committingWizardId, commit, moveSet, salt, otherWizardId, otherCommit)
}

// OneSidedReveal is a paid mutator transaction binding the contract method 0xcef9b488.
//
// Solidity: function oneSidedReveal(uint256 committingWizardId, bytes32 commit, bytes32 moveSet, bytes32 salt, uint256 otherWizardId, bytes32 otherCommit) returns()
func (_BasicTournament *BasicTournamentSession) OneSidedReveal(committingWizardId *big.Int, commit [32]byte, moveSet [32]byte, salt [32]byte, otherWizardId *big.Int, otherCommit [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.OneSidedReveal(&_BasicTournament.TransactOpts, committingWizardId, commit, moveSet, salt, otherWizardId, otherCommit)
}

// OneSidedReveal is a paid mutator transaction binding the contract method 0xcef9b488.
//
// Solidity: function oneSidedReveal(uint256 committingWizardId, bytes32 commit, bytes32 moveSet, bytes32 salt, uint256 otherWizardId, bytes32 otherCommit) returns()
func (_BasicTournament *BasicTournamentTransactorSession) OneSidedReveal(committingWizardId *big.Int, commit [32]byte, moveSet [32]byte, salt [32]byte, otherWizardId *big.Int, otherCommit [32]byte) (*types.Transaction, error) {
	return _BasicTournament.Contract.OneSidedReveal(&_BasicTournament.TransactOpts, committingWizardId, commit, moveSet, salt, otherWizardId, otherCommit)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_BasicTournament *BasicTournamentTransactor) Pause(opts *bind.TransactOpts, pauseDuration *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "pause", pauseDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_BasicTournament *BasicTournamentSession) Pause(pauseDuration *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.Pause(&_BasicTournament.TransactOpts, pauseDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_BasicTournament *BasicTournamentTransactorSession) Pause(pauseDuration *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.Pause(&_BasicTournament.TransactOpts, pauseDuration)
}

// ResolveOneSidedAscensionBattle is a paid mutator transaction binding the contract method 0x35a966f1.
//
// Solidity: function resolveOneSidedAscensionBattle(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) ResolveOneSidedAscensionBattle(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "resolveOneSidedAscensionBattle", wizardId)
}

// ResolveOneSidedAscensionBattle is a paid mutator transaction binding the contract method 0x35a966f1.
//
// Solidity: function resolveOneSidedAscensionBattle(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentSession) ResolveOneSidedAscensionBattle(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ResolveOneSidedAscensionBattle(&_BasicTournament.TransactOpts, wizardId)
}

// ResolveOneSidedAscensionBattle is a paid mutator transaction binding the contract method 0x35a966f1.
//
// Solidity: function resolveOneSidedAscensionBattle(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) ResolveOneSidedAscensionBattle(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ResolveOneSidedAscensionBattle(&_BasicTournament.TransactOpts, wizardId)
}

// ResolveTimedOutDuel is a paid mutator transaction binding the contract method 0xfeb62755.
//
// Solidity: function resolveTimedOutDuel(uint256 wizardId1, uint256 wizardId2) returns()
func (_BasicTournament *BasicTournamentTransactor) ResolveTimedOutDuel(opts *bind.TransactOpts, wizardId1 *big.Int, wizardId2 *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "resolveTimedOutDuel", wizardId1, wizardId2)
}

// ResolveTimedOutDuel is a paid mutator transaction binding the contract method 0xfeb62755.
//
// Solidity: function resolveTimedOutDuel(uint256 wizardId1, uint256 wizardId2) returns()
func (_BasicTournament *BasicTournamentSession) ResolveTimedOutDuel(wizardId1 *big.Int, wizardId2 *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ResolveTimedOutDuel(&_BasicTournament.TransactOpts, wizardId1, wizardId2)
}

// ResolveTimedOutDuel is a paid mutator transaction binding the contract method 0xfeb62755.
//
// Solidity: function resolveTimedOutDuel(uint256 wizardId1, uint256 wizardId2) returns()
func (_BasicTournament *BasicTournamentTransactorSession) ResolveTimedOutDuel(wizardId1 *big.Int, wizardId2 *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.ResolveTimedOutDuel(&_BasicTournament.TransactOpts, wizardId1, wizardId2)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) Revive(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "revive", wizardId)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentSession) Revive(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.Revive(&_BasicTournament.TransactOpts, wizardId)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) Revive(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.Revive(&_BasicTournament.TransactOpts, wizardId)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_BasicTournament *BasicTournamentTransactor) SetCeo(opts *bind.TransactOpts, newCeo common.Address) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "setCeo", newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_BasicTournament *BasicTournamentSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCeo(&_BasicTournament.TransactOpts, newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_BasicTournament *BasicTournamentTransactorSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCeo(&_BasicTournament.TransactOpts, newCeo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_BasicTournament *BasicTournamentTransactor) SetCfo(opts *bind.TransactOpts, newCfo common.Address) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "setCfo", newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_BasicTournament *BasicTournamentSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCfo(&_BasicTournament.TransactOpts, newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_BasicTournament *BasicTournamentTransactorSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCfo(&_BasicTournament.TransactOpts, newCfo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_BasicTournament *BasicTournamentTransactor) SetCoo(opts *bind.TransactOpts, newCoo common.Address) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "setCoo", newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_BasicTournament *BasicTournamentSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCoo(&_BasicTournament.TransactOpts, newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_BasicTournament *BasicTournamentTransactorSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _BasicTournament.Contract.SetCoo(&_BasicTournament.TransactOpts, newCoo)
}

// StartAscension is a paid mutator transaction binding the contract method 0x58042deb.
//
// Solidity: function startAscension(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) StartAscension(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "startAscension", wizardId)
}

// StartAscension is a paid mutator transaction binding the contract method 0x58042deb.
//
// Solidity: function startAscension(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentSession) StartAscension(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.StartAscension(&_BasicTournament.TransactOpts, wizardId)
}

// StartAscension is a paid mutator transaction binding the contract method 0x58042deb.
//
// Solidity: function startAscension(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) StartAscension(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.StartAscension(&_BasicTournament.TransactOpts, wizardId)
}

// UpdateAffinity is a paid mutator transaction binding the contract method 0x5a453d40.
//
// Solidity: function updateAffinity(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactor) UpdateAffinity(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.contract.Transact(opts, "updateAffinity", wizardId)
}

// UpdateAffinity is a paid mutator transaction binding the contract method 0x5a453d40.
//
// Solidity: function updateAffinity(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentSession) UpdateAffinity(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.UpdateAffinity(&_BasicTournament.TransactOpts, wizardId)
}

// UpdateAffinity is a paid mutator transaction binding the contract method 0x5a453d40.
//
// Solidity: function updateAffinity(uint256 wizardId) returns()
func (_BasicTournament *BasicTournamentTransactorSession) UpdateAffinity(wizardId *big.Int) (*types.Transaction, error) {
	return _BasicTournament.Contract.UpdateAffinity(&_BasicTournament.TransactOpts, wizardId)
}

// BasicTournamentAscensionChallengedIterator is returned from FilterAscensionChallenged and is used to iterate over the raw logs and unpacked data for AscensionChallenged events raised by the BasicTournament contract.
type BasicTournamentAscensionChallengedIterator struct {
	Event *BasicTournamentAscensionChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentAscensionChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentAscensionChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentAscensionChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentAscensionChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentAscensionChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentAscensionChallenged represents a AscensionChallenged event raised by the BasicTournament contract.
type BasicTournamentAscensionChallenged struct {
	AscendingWizardId   *big.Int
	ChallengingWizardId *big.Int
	Commitment          [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAscensionChallenged is a free log retrieval operation binding the contract event 0x01b90f216a53259c7a1ed029ef4d7c085b483b5418be4845bcdc933dec110103.
//
// Solidity: event AscensionChallenged(uint256 ascendingWizardId, uint256 challengingWizardId, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) FilterAscensionChallenged(opts *bind.FilterOpts) (*BasicTournamentAscensionChallengedIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "AscensionChallenged")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentAscensionChallengedIterator{contract: _BasicTournament.contract, event: "AscensionChallenged", logs: logs, sub: sub}, nil
}

// WatchAscensionChallenged is a free log subscription operation binding the contract event 0x01b90f216a53259c7a1ed029ef4d7c085b483b5418be4845bcdc933dec110103.
//
// Solidity: event AscensionChallenged(uint256 ascendingWizardId, uint256 challengingWizardId, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) WatchAscensionChallenged(opts *bind.WatchOpts, sink chan<- *BasicTournamentAscensionChallenged) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "AscensionChallenged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentAscensionChallenged)
				if err := _BasicTournament.contract.UnpackLog(event, "AscensionChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAscensionChallenged is a log parse operation binding the contract event 0x01b90f216a53259c7a1ed029ef4d7c085b483b5418be4845bcdc933dec110103.
//
// Solidity: event AscensionChallenged(uint256 ascendingWizardId, uint256 challengingWizardId, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) ParseAscensionChallenged(log types.Log) (*BasicTournamentAscensionChallenged, error) {
	event := new(BasicTournamentAscensionChallenged)
	if err := _BasicTournament.contract.UnpackLog(event, "AscensionChallenged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentAscensionCompleteIterator is returned from FilterAscensionComplete and is used to iterate over the raw logs and unpacked data for AscensionComplete events raised by the BasicTournament contract.
type BasicTournamentAscensionCompleteIterator struct {
	Event *BasicTournamentAscensionComplete // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentAscensionCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentAscensionComplete)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentAscensionComplete)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentAscensionCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentAscensionCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentAscensionComplete represents a AscensionComplete event raised by the BasicTournament contract.
type BasicTournamentAscensionComplete struct {
	WizardId *big.Int
	Power    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAscensionComplete is a free log retrieval operation binding the contract event 0x5b0c8f5b5cdb91a5273ebf46ae7473ad46e9c17fb914ff3e5965ed6a0696c45f.
//
// Solidity: event AscensionComplete(uint256 wizardId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) FilterAscensionComplete(opts *bind.FilterOpts) (*BasicTournamentAscensionCompleteIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "AscensionComplete")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentAscensionCompleteIterator{contract: _BasicTournament.contract, event: "AscensionComplete", logs: logs, sub: sub}, nil
}

// WatchAscensionComplete is a free log subscription operation binding the contract event 0x5b0c8f5b5cdb91a5273ebf46ae7473ad46e9c17fb914ff3e5965ed6a0696c45f.
//
// Solidity: event AscensionComplete(uint256 wizardId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) WatchAscensionComplete(opts *bind.WatchOpts, sink chan<- *BasicTournamentAscensionComplete) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "AscensionComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentAscensionComplete)
				if err := _BasicTournament.contract.UnpackLog(event, "AscensionComplete", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAscensionComplete is a log parse operation binding the contract event 0x5b0c8f5b5cdb91a5273ebf46ae7473ad46e9c17fb914ff3e5965ed6a0696c45f.
//
// Solidity: event AscensionComplete(uint256 wizardId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) ParseAscensionComplete(log types.Log) (*BasicTournamentAscensionComplete, error) {
	event := new(BasicTournamentAscensionComplete)
	if err := _BasicTournament.contract.UnpackLog(event, "AscensionComplete", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentAscensionPairUpIterator is returned from FilterAscensionPairUp and is used to iterate over the raw logs and unpacked data for AscensionPairUp events raised by the BasicTournament contract.
type BasicTournamentAscensionPairUpIterator struct {
	Event *BasicTournamentAscensionPairUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentAscensionPairUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentAscensionPairUp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentAscensionPairUp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentAscensionPairUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentAscensionPairUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentAscensionPairUp represents a AscensionPairUp event raised by the BasicTournament contract.
type BasicTournamentAscensionPairUp struct {
	WizardId1 *big.Int
	WizardId2 *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAscensionPairUp is a free log retrieval operation binding the contract event 0x317afd6701fe06472de1aa6af055861008f7e57c67729e0804c44be9a48facee.
//
// Solidity: event AscensionPairUp(uint256 wizardId1, uint256 wizardId2)
func (_BasicTournament *BasicTournamentFilterer) FilterAscensionPairUp(opts *bind.FilterOpts) (*BasicTournamentAscensionPairUpIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "AscensionPairUp")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentAscensionPairUpIterator{contract: _BasicTournament.contract, event: "AscensionPairUp", logs: logs, sub: sub}, nil
}

// WatchAscensionPairUp is a free log subscription operation binding the contract event 0x317afd6701fe06472de1aa6af055861008f7e57c67729e0804c44be9a48facee.
//
// Solidity: event AscensionPairUp(uint256 wizardId1, uint256 wizardId2)
func (_BasicTournament *BasicTournamentFilterer) WatchAscensionPairUp(opts *bind.WatchOpts, sink chan<- *BasicTournamentAscensionPairUp) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "AscensionPairUp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentAscensionPairUp)
				if err := _BasicTournament.contract.UnpackLog(event, "AscensionPairUp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAscensionPairUp is a log parse operation binding the contract event 0x317afd6701fe06472de1aa6af055861008f7e57c67729e0804c44be9a48facee.
//
// Solidity: event AscensionPairUp(uint256 wizardId1, uint256 wizardId2)
func (_BasicTournament *BasicTournamentFilterer) ParseAscensionPairUp(log types.Log) (*BasicTournamentAscensionPairUp, error) {
	event := new(BasicTournamentAscensionPairUp)
	if err := _BasicTournament.contract.UnpackLog(event, "AscensionPairUp", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentAscensionStartIterator is returned from FilterAscensionStart and is used to iterate over the raw logs and unpacked data for AscensionStart events raised by the BasicTournament contract.
type BasicTournamentAscensionStartIterator struct {
	Event *BasicTournamentAscensionStart // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentAscensionStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentAscensionStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentAscensionStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentAscensionStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentAscensionStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentAscensionStart represents a AscensionStart event raised by the BasicTournament contract.
type BasicTournamentAscensionStart struct {
	WizardId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAscensionStart is a free log retrieval operation binding the contract event 0x75e589c241c6adb443794af5ecb220cdeb59eb1c44a37c6d1dfaebde22706b51.
//
// Solidity: event AscensionStart(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) FilterAscensionStart(opts *bind.FilterOpts) (*BasicTournamentAscensionStartIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "AscensionStart")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentAscensionStartIterator{contract: _BasicTournament.contract, event: "AscensionStart", logs: logs, sub: sub}, nil
}

// WatchAscensionStart is a free log subscription operation binding the contract event 0x75e589c241c6adb443794af5ecb220cdeb59eb1c44a37c6d1dfaebde22706b51.
//
// Solidity: event AscensionStart(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) WatchAscensionStart(opts *bind.WatchOpts, sink chan<- *BasicTournamentAscensionStart) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "AscensionStart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentAscensionStart)
				if err := _BasicTournament.contract.UnpackLog(event, "AscensionStart", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAscensionStart is a log parse operation binding the contract event 0x75e589c241c6adb443794af5ecb220cdeb59eb1c44a37c6d1dfaebde22706b51.
//
// Solidity: event AscensionStart(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) ParseAscensionStart(log types.Log) (*BasicTournamentAscensionStart, error) {
	event := new(BasicTournamentAscensionStart)
	if err := _BasicTournament.contract.UnpackLog(event, "AscensionStart", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentCEOTransferredIterator is returned from FilterCEOTransferred and is used to iterate over the raw logs and unpacked data for CEOTransferred events raised by the BasicTournament contract.
type BasicTournamentCEOTransferredIterator struct {
	Event *BasicTournamentCEOTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentCEOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentCEOTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentCEOTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentCEOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentCEOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentCEOTransferred represents a CEOTransferred event raised by the BasicTournament contract.
type BasicTournamentCEOTransferred struct {
	PreviousCeo common.Address
	NewCeo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCEOTransferred is a free log retrieval operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_BasicTournament *BasicTournamentFilterer) FilterCEOTransferred(opts *bind.FilterOpts) (*BasicTournamentCEOTransferredIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentCEOTransferredIterator{contract: _BasicTournament.contract, event: "CEOTransferred", logs: logs, sub: sub}, nil
}

// WatchCEOTransferred is a free log subscription operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_BasicTournament *BasicTournamentFilterer) WatchCEOTransferred(opts *bind.WatchOpts, sink chan<- *BasicTournamentCEOTransferred) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentCEOTransferred)
				if err := _BasicTournament.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCEOTransferred is a log parse operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_BasicTournament *BasicTournamentFilterer) ParseCEOTransferred(log types.Log) (*BasicTournamentCEOTransferred, error) {
	event := new(BasicTournamentCEOTransferred)
	if err := _BasicTournament.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentCFOTransferredIterator is returned from FilterCFOTransferred and is used to iterate over the raw logs and unpacked data for CFOTransferred events raised by the BasicTournament contract.
type BasicTournamentCFOTransferredIterator struct {
	Event *BasicTournamentCFOTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentCFOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentCFOTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentCFOTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentCFOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentCFOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentCFOTransferred represents a CFOTransferred event raised by the BasicTournament contract.
type BasicTournamentCFOTransferred struct {
	PreviousCfo common.Address
	NewCfo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCFOTransferred is a free log retrieval operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_BasicTournament *BasicTournamentFilterer) FilterCFOTransferred(opts *bind.FilterOpts) (*BasicTournamentCFOTransferredIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentCFOTransferredIterator{contract: _BasicTournament.contract, event: "CFOTransferred", logs: logs, sub: sub}, nil
}

// WatchCFOTransferred is a free log subscription operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_BasicTournament *BasicTournamentFilterer) WatchCFOTransferred(opts *bind.WatchOpts, sink chan<- *BasicTournamentCFOTransferred) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentCFOTransferred)
				if err := _BasicTournament.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCFOTransferred is a log parse operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_BasicTournament *BasicTournamentFilterer) ParseCFOTransferred(log types.Log) (*BasicTournamentCFOTransferred, error) {
	event := new(BasicTournamentCFOTransferred)
	if err := _BasicTournament.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentCOOTransferredIterator is returned from FilterCOOTransferred and is used to iterate over the raw logs and unpacked data for COOTransferred events raised by the BasicTournament contract.
type BasicTournamentCOOTransferredIterator struct {
	Event *BasicTournamentCOOTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentCOOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentCOOTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentCOOTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentCOOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentCOOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentCOOTransferred represents a COOTransferred event raised by the BasicTournament contract.
type BasicTournamentCOOTransferred struct {
	PreviousCoo common.Address
	NewCoo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCOOTransferred is a free log retrieval operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_BasicTournament *BasicTournamentFilterer) FilterCOOTransferred(opts *bind.FilterOpts) (*BasicTournamentCOOTransferredIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentCOOTransferredIterator{contract: _BasicTournament.contract, event: "COOTransferred", logs: logs, sub: sub}, nil
}

// WatchCOOTransferred is a free log subscription operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_BasicTournament *BasicTournamentFilterer) WatchCOOTransferred(opts *bind.WatchOpts, sink chan<- *BasicTournamentCOOTransferred) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentCOOTransferred)
				if err := _BasicTournament.contract.UnpackLog(event, "COOTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCOOTransferred is a log parse operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_BasicTournament *BasicTournamentFilterer) ParseCOOTransferred(log types.Log) (*BasicTournamentCOOTransferred, error) {
	event := new(BasicTournamentCOOTransferred)
	if err := _BasicTournament.contract.UnpackLog(event, "COOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentDuelEndIterator is returned from FilterDuelEnd and is used to iterate over the raw logs and unpacked data for DuelEnd events raised by the BasicTournament contract.
type BasicTournamentDuelEndIterator struct {
	Event *BasicTournamentDuelEnd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentDuelEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentDuelEnd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentDuelEnd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentDuelEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentDuelEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentDuelEnd represents a DuelEnd event raised by the BasicTournament contract.
type BasicTournamentDuelEnd struct {
	DuelId    [32]byte
	WizardId1 *big.Int
	WizardId2 *big.Int
	MoveSet1  [32]byte
	MoveSet2  [32]byte
	Power1    *big.Int
	Power2    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDuelEnd is a free log retrieval operation binding the contract event 0x5d340f045b88c840424b9dc0322566ac25ac52a1710233de27fec1fcd04eae60.
//
// Solidity: event DuelEnd(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) FilterDuelEnd(opts *bind.FilterOpts) (*BasicTournamentDuelEndIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "DuelEnd")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentDuelEndIterator{contract: _BasicTournament.contract, event: "DuelEnd", logs: logs, sub: sub}, nil
}

// WatchDuelEnd is a free log subscription operation binding the contract event 0x5d340f045b88c840424b9dc0322566ac25ac52a1710233de27fec1fcd04eae60.
//
// Solidity: event DuelEnd(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) WatchDuelEnd(opts *bind.WatchOpts, sink chan<- *BasicTournamentDuelEnd) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "DuelEnd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentDuelEnd)
				if err := _BasicTournament.contract.UnpackLog(event, "DuelEnd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDuelEnd is a log parse operation binding the contract event 0x5d340f045b88c840424b9dc0322566ac25ac52a1710233de27fec1fcd04eae60.
//
// Solidity: event DuelEnd(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) ParseDuelEnd(log types.Log) (*BasicTournamentDuelEnd, error) {
	event := new(BasicTournamentDuelEnd)
	if err := _BasicTournament.contract.UnpackLog(event, "DuelEnd", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentDuelStartIterator is returned from FilterDuelStart and is used to iterate over the raw logs and unpacked data for DuelStart events raised by the BasicTournament contract.
type BasicTournamentDuelStartIterator struct {
	Event *BasicTournamentDuelStart // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentDuelStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentDuelStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentDuelStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentDuelStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentDuelStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentDuelStart represents a DuelStart event raised by the BasicTournament contract.
type BasicTournamentDuelStart struct {
	DuelId            [32]byte
	WizardId1         *big.Int
	WizardId2         *big.Int
	TimeoutBlock      *big.Int
	IsAscensionBattle bool
	Commit1           [32]byte
	Commit2           [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDuelStart is a free log retrieval operation binding the contract event 0x77d9948186effe9235713be25320cbda824d0fac31841e1af13e014cb0b34853.
//
// Solidity: event DuelStart(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 timeoutBlock, bool isAscensionBattle, bytes32 commit1, bytes32 commit2)
func (_BasicTournament *BasicTournamentFilterer) FilterDuelStart(opts *bind.FilterOpts) (*BasicTournamentDuelStartIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "DuelStart")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentDuelStartIterator{contract: _BasicTournament.contract, event: "DuelStart", logs: logs, sub: sub}, nil
}

// WatchDuelStart is a free log subscription operation binding the contract event 0x77d9948186effe9235713be25320cbda824d0fac31841e1af13e014cb0b34853.
//
// Solidity: event DuelStart(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 timeoutBlock, bool isAscensionBattle, bytes32 commit1, bytes32 commit2)
func (_BasicTournament *BasicTournamentFilterer) WatchDuelStart(opts *bind.WatchOpts, sink chan<- *BasicTournamentDuelStart) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "DuelStart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentDuelStart)
				if err := _BasicTournament.contract.UnpackLog(event, "DuelStart", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDuelStart is a log parse operation binding the contract event 0x77d9948186effe9235713be25320cbda824d0fac31841e1af13e014cb0b34853.
//
// Solidity: event DuelStart(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 timeoutBlock, bool isAscensionBattle, bytes32 commit1, bytes32 commit2)
func (_BasicTournament *BasicTournamentFilterer) ParseDuelStart(log types.Log) (*BasicTournamentDuelStart, error) {
	event := new(BasicTournamentDuelStart)
	if err := _BasicTournament.contract.UnpackLog(event, "DuelStart", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentDuelTimeOutIterator is returned from FilterDuelTimeOut and is used to iterate over the raw logs and unpacked data for DuelTimeOut events raised by the BasicTournament contract.
type BasicTournamentDuelTimeOutIterator struct {
	Event *BasicTournamentDuelTimeOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentDuelTimeOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentDuelTimeOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentDuelTimeOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentDuelTimeOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentDuelTimeOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentDuelTimeOut represents a DuelTimeOut event raised by the BasicTournament contract.
type BasicTournamentDuelTimeOut struct {
	DuelId    [32]byte
	WizardId1 *big.Int
	WizardId2 *big.Int
	Power1    *big.Int
	Power2    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDuelTimeOut is a free log retrieval operation binding the contract event 0xc7aa367585942c098842dd2573f26d69dabb55248bcaea3c4463cd96680e1c2e.
//
// Solidity: event DuelTimeOut(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) FilterDuelTimeOut(opts *bind.FilterOpts) (*BasicTournamentDuelTimeOutIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "DuelTimeOut")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentDuelTimeOutIterator{contract: _BasicTournament.contract, event: "DuelTimeOut", logs: logs, sub: sub}, nil
}

// WatchDuelTimeOut is a free log subscription operation binding the contract event 0xc7aa367585942c098842dd2573f26d69dabb55248bcaea3c4463cd96680e1c2e.
//
// Solidity: event DuelTimeOut(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) WatchDuelTimeOut(opts *bind.WatchOpts, sink chan<- *BasicTournamentDuelTimeOut) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "DuelTimeOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentDuelTimeOut)
				if err := _BasicTournament.contract.UnpackLog(event, "DuelTimeOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDuelTimeOut is a log parse operation binding the contract event 0xc7aa367585942c098842dd2573f26d69dabb55248bcaea3c4463cd96680e1c2e.
//
// Solidity: event DuelTimeOut(bytes32 duelId, uint256 wizardId1, uint256 wizardId2, uint256 power1, uint256 power2)
func (_BasicTournament *BasicTournamentFilterer) ParseDuelTimeOut(log types.Log) (*BasicTournamentDuelTimeOut, error) {
	event := new(BasicTournamentDuelTimeOut)
	if err := _BasicTournament.contract.UnpackLog(event, "DuelTimeOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentOneSidedCommitAddedIterator is returned from FilterOneSidedCommitAdded and is used to iterate over the raw logs and unpacked data for OneSidedCommitAdded events raised by the BasicTournament contract.
type BasicTournamentOneSidedCommitAddedIterator struct {
	Event *BasicTournamentOneSidedCommitAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentOneSidedCommitAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentOneSidedCommitAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentOneSidedCommitAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentOneSidedCommitAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentOneSidedCommitAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentOneSidedCommitAdded represents a OneSidedCommitAdded event raised by the BasicTournament contract.
type BasicTournamentOneSidedCommitAdded struct {
	CommittingWizardId    *big.Int
	OtherWizardId         *big.Int
	CommittingWizardNonce *big.Int
	OtherWizardNonce      *big.Int
	Commitment            [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOneSidedCommitAdded is a free log retrieval operation binding the contract event 0x530a6602289d4bdb1c24e67332be34dd86bcf90f97911cf9ab69cf9d48db5eeb.
//
// Solidity: event OneSidedCommitAdded(uint256 committingWizardId, uint256 otherWizardId, uint256 committingWizardNonce, uint256 otherWizardNonce, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) FilterOneSidedCommitAdded(opts *bind.FilterOpts) (*BasicTournamentOneSidedCommitAddedIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "OneSidedCommitAdded")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentOneSidedCommitAddedIterator{contract: _BasicTournament.contract, event: "OneSidedCommitAdded", logs: logs, sub: sub}, nil
}

// WatchOneSidedCommitAdded is a free log subscription operation binding the contract event 0x530a6602289d4bdb1c24e67332be34dd86bcf90f97911cf9ab69cf9d48db5eeb.
//
// Solidity: event OneSidedCommitAdded(uint256 committingWizardId, uint256 otherWizardId, uint256 committingWizardNonce, uint256 otherWizardNonce, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) WatchOneSidedCommitAdded(opts *bind.WatchOpts, sink chan<- *BasicTournamentOneSidedCommitAdded) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "OneSidedCommitAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentOneSidedCommitAdded)
				if err := _BasicTournament.contract.UnpackLog(event, "OneSidedCommitAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneSidedCommitAdded is a log parse operation binding the contract event 0x530a6602289d4bdb1c24e67332be34dd86bcf90f97911cf9ab69cf9d48db5eeb.
//
// Solidity: event OneSidedCommitAdded(uint256 committingWizardId, uint256 otherWizardId, uint256 committingWizardNonce, uint256 otherWizardNonce, bytes32 commitment)
func (_BasicTournament *BasicTournamentFilterer) ParseOneSidedCommitAdded(log types.Log) (*BasicTournamentOneSidedCommitAdded, error) {
	event := new(BasicTournamentOneSidedCommitAdded)
	if err := _BasicTournament.contract.UnpackLog(event, "OneSidedCommitAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentOneSidedCommitCancelledIterator is returned from FilterOneSidedCommitCancelled and is used to iterate over the raw logs and unpacked data for OneSidedCommitCancelled events raised by the BasicTournament contract.
type BasicTournamentOneSidedCommitCancelledIterator struct {
	Event *BasicTournamentOneSidedCommitCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentOneSidedCommitCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentOneSidedCommitCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentOneSidedCommitCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentOneSidedCommitCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentOneSidedCommitCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentOneSidedCommitCancelled represents a OneSidedCommitCancelled event raised by the BasicTournament contract.
type BasicTournamentOneSidedCommitCancelled struct {
	WizardId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneSidedCommitCancelled is a free log retrieval operation binding the contract event 0x04d33264c6183366c5711d870e69908ac063f2c36c434c4b1f77916c238b9dfc.
//
// Solidity: event OneSidedCommitCancelled(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) FilterOneSidedCommitCancelled(opts *bind.FilterOpts) (*BasicTournamentOneSidedCommitCancelledIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "OneSidedCommitCancelled")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentOneSidedCommitCancelledIterator{contract: _BasicTournament.contract, event: "OneSidedCommitCancelled", logs: logs, sub: sub}, nil
}

// WatchOneSidedCommitCancelled is a free log subscription operation binding the contract event 0x04d33264c6183366c5711d870e69908ac063f2c36c434c4b1f77916c238b9dfc.
//
// Solidity: event OneSidedCommitCancelled(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) WatchOneSidedCommitCancelled(opts *bind.WatchOpts, sink chan<- *BasicTournamentOneSidedCommitCancelled) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "OneSidedCommitCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentOneSidedCommitCancelled)
				if err := _BasicTournament.contract.UnpackLog(event, "OneSidedCommitCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneSidedCommitCancelled is a log parse operation binding the contract event 0x04d33264c6183366c5711d870e69908ac063f2c36c434c4b1f77916c238b9dfc.
//
// Solidity: event OneSidedCommitCancelled(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) ParseOneSidedCommitCancelled(log types.Log) (*BasicTournamentOneSidedCommitCancelled, error) {
	event := new(BasicTournamentOneSidedCommitCancelled)
	if err := _BasicTournament.contract.UnpackLog(event, "OneSidedCommitCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentOneSidedRevealAddedIterator is returned from FilterOneSidedRevealAdded and is used to iterate over the raw logs and unpacked data for OneSidedRevealAdded events raised by the BasicTournament contract.
type BasicTournamentOneSidedRevealAddedIterator struct {
	Event *BasicTournamentOneSidedRevealAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentOneSidedRevealAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentOneSidedRevealAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentOneSidedRevealAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentOneSidedRevealAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentOneSidedRevealAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentOneSidedRevealAdded represents a OneSidedRevealAdded event raised by the BasicTournament contract.
type BasicTournamentOneSidedRevealAdded struct {
	DuelId             [32]byte
	CommittingWizardId *big.Int
	OtherWizardId      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOneSidedRevealAdded is a free log retrieval operation binding the contract event 0x4761d35fc455d658f13ab27c916c9772625a2cd266c1e36b1e03b6bbe9e03f1a.
//
// Solidity: event OneSidedRevealAdded(bytes32 duelId, uint256 committingWizardId, uint256 otherWizardId)
func (_BasicTournament *BasicTournamentFilterer) FilterOneSidedRevealAdded(opts *bind.FilterOpts) (*BasicTournamentOneSidedRevealAddedIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "OneSidedRevealAdded")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentOneSidedRevealAddedIterator{contract: _BasicTournament.contract, event: "OneSidedRevealAdded", logs: logs, sub: sub}, nil
}

// WatchOneSidedRevealAdded is a free log subscription operation binding the contract event 0x4761d35fc455d658f13ab27c916c9772625a2cd266c1e36b1e03b6bbe9e03f1a.
//
// Solidity: event OneSidedRevealAdded(bytes32 duelId, uint256 committingWizardId, uint256 otherWizardId)
func (_BasicTournament *BasicTournamentFilterer) WatchOneSidedRevealAdded(opts *bind.WatchOpts, sink chan<- *BasicTournamentOneSidedRevealAdded) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "OneSidedRevealAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentOneSidedRevealAdded)
				if err := _BasicTournament.contract.UnpackLog(event, "OneSidedRevealAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneSidedRevealAdded is a log parse operation binding the contract event 0x4761d35fc455d658f13ab27c916c9772625a2cd266c1e36b1e03b6bbe9e03f1a.
//
// Solidity: event OneSidedRevealAdded(bytes32 duelId, uint256 committingWizardId, uint256 otherWizardId)
func (_BasicTournament *BasicTournamentFilterer) ParseOneSidedRevealAdded(log types.Log) (*BasicTournamentOneSidedRevealAdded, error) {
	event := new(BasicTournamentOneSidedRevealAdded)
	if err := _BasicTournament.contract.UnpackLog(event, "OneSidedRevealAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BasicTournament contract.
type BasicTournamentPausedIterator struct {
	Event *BasicTournamentPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentPaused represents a Paused event raised by the BasicTournament contract.
type BasicTournamentPaused struct {
	PauseEndedBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_BasicTournament *BasicTournamentFilterer) FilterPaused(opts *bind.FilterOpts) (*BasicTournamentPausedIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentPausedIterator{contract: _BasicTournament.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_BasicTournament *BasicTournamentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BasicTournamentPaused) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentPaused)
				if err := _BasicTournament.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_BasicTournament *BasicTournamentFilterer) ParsePaused(log types.Log) (*BasicTournamentPaused, error) {
	event := new(BasicTournamentPaused)
	if err := _BasicTournament.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentPowerTransferredIterator is returned from FilterPowerTransferred and is used to iterate over the raw logs and unpacked data for PowerTransferred events raised by the BasicTournament contract.
type BasicTournamentPowerTransferredIterator struct {
	Event *BasicTournamentPowerTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentPowerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentPowerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentPowerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentPowerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentPowerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentPowerTransferred represents a PowerTransferred event raised by the BasicTournament contract.
type BasicTournamentPowerTransferred struct {
	SendingWizId      *big.Int
	ReceivingWizId    *big.Int
	AmountTransferred *big.Int
	Reason            uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPowerTransferred is a free log retrieval operation binding the contract event 0x6f76f2ae8cea65409255207333bfb6f8c087e8e3b6b3f8cc7798fc35ed7af99a.
//
// Solidity: event PowerTransferred(uint256 sendingWizId, uint256 receivingWizId, uint256 amountTransferred, uint8 reason)
func (_BasicTournament *BasicTournamentFilterer) FilterPowerTransferred(opts *bind.FilterOpts) (*BasicTournamentPowerTransferredIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "PowerTransferred")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentPowerTransferredIterator{contract: _BasicTournament.contract, event: "PowerTransferred", logs: logs, sub: sub}, nil
}

// WatchPowerTransferred is a free log subscription operation binding the contract event 0x6f76f2ae8cea65409255207333bfb6f8c087e8e3b6b3f8cc7798fc35ed7af99a.
//
// Solidity: event PowerTransferred(uint256 sendingWizId, uint256 receivingWizId, uint256 amountTransferred, uint8 reason)
func (_BasicTournament *BasicTournamentFilterer) WatchPowerTransferred(opts *bind.WatchOpts, sink chan<- *BasicTournamentPowerTransferred) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "PowerTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentPowerTransferred)
				if err := _BasicTournament.contract.UnpackLog(event, "PowerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePowerTransferred is a log parse operation binding the contract event 0x6f76f2ae8cea65409255207333bfb6f8c087e8e3b6b3f8cc7798fc35ed7af99a.
//
// Solidity: event PowerTransferred(uint256 sendingWizId, uint256 receivingWizId, uint256 amountTransferred, uint8 reason)
func (_BasicTournament *BasicTournamentFilterer) ParsePowerTransferred(log types.Log) (*BasicTournamentPowerTransferred, error) {
	event := new(BasicTournamentPowerTransferred)
	if err := _BasicTournament.contract.UnpackLog(event, "PowerTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentPrizeClaimedIterator is returned from FilterPrizeClaimed and is used to iterate over the raw logs and unpacked data for PrizeClaimed events raised by the BasicTournament contract.
type BasicTournamentPrizeClaimedIterator struct {
	Event *BasicTournamentPrizeClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentPrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentPrizeClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentPrizeClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentPrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentPrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentPrizeClaimed represents a PrizeClaimed event raised by the BasicTournament contract.
type BasicTournamentPrizeClaimed struct {
	ClaimingWinnerId *big.Int
	PrizeAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimed is a free log retrieval operation binding the contract event 0xd53b67ba94a5d6268d11caa5d2693557779404ed02fc9825d86d2894d29cb8fd.
//
// Solidity: event PrizeClaimed(uint256 claimingWinnerId, uint256 prizeAmount)
func (_BasicTournament *BasicTournamentFilterer) FilterPrizeClaimed(opts *bind.FilterOpts) (*BasicTournamentPrizeClaimedIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentPrizeClaimedIterator{contract: _BasicTournament.contract, event: "PrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimed is a free log subscription operation binding the contract event 0xd53b67ba94a5d6268d11caa5d2693557779404ed02fc9825d86d2894d29cb8fd.
//
// Solidity: event PrizeClaimed(uint256 claimingWinnerId, uint256 prizeAmount)
func (_BasicTournament *BasicTournamentFilterer) WatchPrizeClaimed(opts *bind.WatchOpts, sink chan<- *BasicTournamentPrizeClaimed) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentPrizeClaimed)
				if err := _BasicTournament.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePrizeClaimed is a log parse operation binding the contract event 0xd53b67ba94a5d6268d11caa5d2693557779404ed02fc9825d86d2894d29cb8fd.
//
// Solidity: event PrizeClaimed(uint256 claimingWinnerId, uint256 prizeAmount)
func (_BasicTournament *BasicTournamentFilterer) ParsePrizeClaimed(log types.Log) (*BasicTournamentPrizeClaimed, error) {
	event := new(BasicTournamentPrizeClaimed)
	if err := _BasicTournament.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentReviveIterator is returned from FilterRevive and is used to iterate over the raw logs and unpacked data for Revive events raised by the BasicTournament contract.
type BasicTournamentReviveIterator struct {
	Event *BasicTournamentRevive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentReviveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentRevive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentRevive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentReviveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentReviveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentRevive represents a Revive event raised by the BasicTournament contract.
type BasicTournamentRevive struct {
	WizId *big.Int
	Power *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevive is a free log retrieval operation binding the contract event 0xa78677222d515efffcb323b960622c3e2bff3331916798f375b592c2a07f6c5a.
//
// Solidity: event Revive(uint256 wizId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) FilterRevive(opts *bind.FilterOpts) (*BasicTournamentReviveIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "Revive")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentReviveIterator{contract: _BasicTournament.contract, event: "Revive", logs: logs, sub: sub}, nil
}

// WatchRevive is a free log subscription operation binding the contract event 0xa78677222d515efffcb323b960622c3e2bff3331916798f375b592c2a07f6c5a.
//
// Solidity: event Revive(uint256 wizId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) WatchRevive(opts *bind.WatchOpts, sink chan<- *BasicTournamentRevive) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "Revive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentRevive)
				if err := _BasicTournament.contract.UnpackLog(event, "Revive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevive is a log parse operation binding the contract event 0xa78677222d515efffcb323b960622c3e2bff3331916798f375b592c2a07f6c5a.
//
// Solidity: event Revive(uint256 wizId, uint256 power)
func (_BasicTournament *BasicTournamentFilterer) ParseRevive(log types.Log) (*BasicTournamentRevive, error) {
	event := new(BasicTournamentRevive)
	if err := _BasicTournament.contract.UnpackLog(event, "Revive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BasicTournamentWizardEliminationIterator is returned from FilterWizardElimination and is used to iterate over the raw logs and unpacked data for WizardElimination events raised by the BasicTournament contract.
type BasicTournamentWizardEliminationIterator struct {
	Event *BasicTournamentWizardElimination // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTournamentWizardEliminationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTournamentWizardElimination)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTournamentWizardElimination)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTournamentWizardEliminationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTournamentWizardEliminationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTournamentWizardElimination represents a WizardElimination event raised by the BasicTournament contract.
type BasicTournamentWizardElimination struct {
	WizardId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWizardElimination is a free log retrieval operation binding the contract event 0x467c36a9931be143929e59fbcf25cc5ea17577e78af28da39a7b5f23d9c081b3.
//
// Solidity: event WizardElimination(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) FilterWizardElimination(opts *bind.FilterOpts) (*BasicTournamentWizardEliminationIterator, error) {

	logs, sub, err := _BasicTournament.contract.FilterLogs(opts, "WizardElimination")
	if err != nil {
		return nil, err
	}
	return &BasicTournamentWizardEliminationIterator{contract: _BasicTournament.contract, event: "WizardElimination", logs: logs, sub: sub}, nil
}

// WatchWizardElimination is a free log subscription operation binding the contract event 0x467c36a9931be143929e59fbcf25cc5ea17577e78af28da39a7b5f23d9c081b3.
//
// Solidity: event WizardElimination(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) WatchWizardElimination(opts *bind.WatchOpts, sink chan<- *BasicTournamentWizardElimination) (event.Subscription, error) {

	logs, sub, err := _BasicTournament.contract.WatchLogs(opts, "WizardElimination")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTournamentWizardElimination)
				if err := _BasicTournament.contract.UnpackLog(event, "WizardElimination", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWizardElimination is a log parse operation binding the contract event 0x467c36a9931be143929e59fbcf25cc5ea17577e78af28da39a7b5f23d9c081b3.
//
// Solidity: event WizardElimination(uint256 wizardId)
func (_BasicTournament *BasicTournamentFilterer) ParseWizardElimination(log types.Log) (*BasicTournamentWizardElimination, error) {
	event := new(BasicTournamentWizardElimination)
	if err := _BasicTournament.contract.UnpackLog(event, "WizardElimination", log); err != nil {
		return nil, err
	}
	return event, nil
}
