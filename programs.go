package solana

import "github.com/near/borsh-go"

var (
	// Native programs
	SystemProgram             Pubkey = MustParsePubkey("11111111111111111111111111111111")            //Create new accounts, allocate account data, assign accounts to owning programs, transfer lamports from System Program owned accounts and pay transaction fees.
	ConfigProgram             Pubkey = MustParsePubkey("Config1111111111111111111111111111111111111") //Add configuration data to the chain, followed by the list of public keys that are allowed to modify it
	StakeProgram              Pubkey = MustParsePubkey("Stake11111111111111111111111111111111111111") //Create and manage accounts representing stake and rewards for delegations to validators.
	VoteProgram               Pubkey = MustParsePubkey("Vote111111111111111111111111111111111111111") //Create and manage accounts that track validator voting state and rewards.
	AddressLookupTableProgram Pubkey = MustParsePubkey("AddressLookupTab1e1111111111111111111111111")
	BpfLoaderProgram          Pubkey = MustParsePubkey("BPFLoaderUpgradeab1e11111111111111111111111") //Deploys, upgrades, and executes programs on the chain.
	Ed25519Program            Pubkey = MustParsePubkey("Ed25519SigVerify111111111111111111111111111") //The program for verifying ed25519 signatures. It takes an ed25519 signature, a public key, and a message. Multiple signatures can be verified. If any of the signatures fail to verify, an error is returned.
	Secp256k1Program          Pubkey = MustParsePubkey("KeccakSecp256k11111111111111111111111111111") //Verify secp256k1 public key recovery operations (ecrecover).
	Secp256r1Program          Pubkey = MustParsePubkey("Secp256r1SigVerify1111111111111111111111111") //The program for verifying secp256r1 signatures. It takes a secp256r1 signature, a public key, and a message. Up to 8 signatures can be verified. If any of the signatures fail to verify, an error is returned.
)

type SystemProgramIxs interface {
	// CreateAccount(fundingAccount Pubkey, newAccount Pubkey, lamports uint, space uint, programOwner Pubkey) Instruction
	// Assign(assignedAccount Pubkey, programOwner Pubkey) Instruction
	Transfer(source Pubkey, destination Pubkey, lamports uint) Instruction
	// CreateAccountWithSeed(fundingAccount Pubkey, createdAccount Pubkey, base Pubkey, seed string, lamports uint, space uint, programOwner Pubkey) Instruction
	// AdvanceNonceAccount(nonceAccount Pubkey, nonceAuthority Pubkey) Instruction
	// WithdrawNonceAccount(nonceAccount Pubkey, destination Pubkey, nonceAuthority Pubkey, lamports uint) Instruction
	// InitializeNonceAccount(nonceAccount Pubkey, nonceAuthority Pubkey) Instruction
	// AuthorizeNonceAccount(nonceAccount Pubkey, nonceAuthority Pubkey, newAuthority Pubkey) Instruction
	// Allocate(allocatedAccount Pubkey, space uint) Instruction
	// AllocateWithSeed(allocatedAccount Pubkey, base Pubkey, seed string, space uint, programOwner Pubkey) Instruction
	// AssignWithSeed(assignedAccount Pubkey, base Pubkey, seed string, programOwner Pubkey) Instruction
	// TransferWithSeed(source Pubkey, base Pubkey, destination Pubkey, lamports uint, fromSeed string, fromOwner Pubkey) Instruction
	// UpgradeNonceAccount(nonceAccount Pubkey, nonceAuthority Pubkey) Instruction
}

func SystemProgramInstructions() SystemProgramIxs {
	return &systemProgramIxs{}
}

type systemProgramIxs struct{}

func (systemProgramIxs) Transfer(source Pubkey, destination Pubkey, lamports uint) Instruction {
	data, _ := borsh.Serialize(struct {
		Instruction uint8
		Lamports    uint64
	}{
		Instruction: 2,
		Lamports:    uint64(lamports),
	})
	return Instruction{
		ProgramID: SystemProgram,
		Data:      data,
		Accounts:  []AccountMeta{{Pubkey: source, Signer: true, Writable: true}, {Pubkey: destination, Signer: false, Writable: true}},
	}
}
