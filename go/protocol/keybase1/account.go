// Auto-generated by avdl-compiler v1.3.21 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/account.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type HasServerKeysRes struct {
	HasServerKeys bool `codec:"hasServerKeys" json:"hasServerKeys"`
}

func (o HasServerKeysRes) DeepCopy() HasServerKeysRes {
	return HasServerKeysRes{
		HasServerKeys: o.HasServerKeys,
	}
}

type PassphraseChangeArg struct {
	SessionID     int    `codec:"sessionID" json:"sessionID"`
	OldPassphrase string `codec:"oldPassphrase" json:"oldPassphrase"`
	Passphrase    string `codec:"passphrase" json:"passphrase"`
	Force         bool   `codec:"force" json:"force"`
}

type PassphrasePromptArg struct {
	SessionID int         `codec:"sessionID" json:"sessionID"`
	GuiArg    GUIEntryArg `codec:"guiArg" json:"guiArg"`
}

type EmailChangeArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	NewEmail  string `codec:"newEmail" json:"newEmail"`
}

type HasServerKeysArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ResetAccountArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type AccountInterface interface {
	// Change the passphrase from old to new. If old isn't set, and force is false,
	// then prompt at the UI for it. If old isn't set and force is true, then we'll
	// try to force a passphrase change.
	PassphraseChange(context.Context, PassphraseChangeArg) error
	PassphrasePrompt(context.Context, PassphrasePromptArg) (GetPassphraseRes, error)
	// * change email to the new given email by signing a statement.
	EmailChange(context.Context, EmailChangeArg) error
	// * Whether the logged-in user has uploaded private keys
	// * Will error if not logged in.
	HasServerKeys(context.Context, int) (HasServerKeysRes, error)
	ResetAccount(context.Context, int) error
}

func AccountProtocol(i AccountInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.account",
		Methods: map[string]rpc.ServeHandlerDescription{
			"passphraseChange": {
				MakeArg: func() interface{} {
					ret := make([]PassphraseChangeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PassphraseChangeArg)
					if !ok {
						err = rpc.NewTypeError((*[]PassphraseChangeArg)(nil), args)
						return
					}
					err = i.PassphraseChange(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"passphrasePrompt": {
				MakeArg: func() interface{} {
					ret := make([]PassphrasePromptArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PassphrasePromptArg)
					if !ok {
						err = rpc.NewTypeError((*[]PassphrasePromptArg)(nil), args)
						return
					}
					ret, err = i.PassphrasePrompt(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"emailChange": {
				MakeArg: func() interface{} {
					ret := make([]EmailChangeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]EmailChangeArg)
					if !ok {
						err = rpc.NewTypeError((*[]EmailChangeArg)(nil), args)
						return
					}
					err = i.EmailChange(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"hasServerKeys": {
				MakeArg: func() interface{} {
					ret := make([]HasServerKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]HasServerKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]HasServerKeysArg)(nil), args)
						return
					}
					ret, err = i.HasServerKeys(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"resetAccount": {
				MakeArg: func() interface{} {
					ret := make([]ResetAccountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ResetAccountArg)
					if !ok {
						err = rpc.NewTypeError((*[]ResetAccountArg)(nil), args)
						return
					}
					err = i.ResetAccount(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type AccountClient struct {
	Cli rpc.GenericClient
}

// Change the passphrase from old to new. If old isn't set, and force is false,
// then prompt at the UI for it. If old isn't set and force is true, then we'll
// try to force a passphrase change.
func (c AccountClient) PassphraseChange(ctx context.Context, __arg PassphraseChangeArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.account.passphraseChange", []interface{}{__arg}, nil)
	return
}

func (c AccountClient) PassphrasePrompt(ctx context.Context, __arg PassphrasePromptArg) (res GetPassphraseRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.account.passphrasePrompt", []interface{}{__arg}, &res)
	return
}

// * change email to the new given email by signing a statement.
func (c AccountClient) EmailChange(ctx context.Context, __arg EmailChangeArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.account.emailChange", []interface{}{__arg}, nil)
	return
}

// * Whether the logged-in user has uploaded private keys
// * Will error if not logged in.
func (c AccountClient) HasServerKeys(ctx context.Context, sessionID int) (res HasServerKeysRes, err error) {
	__arg := HasServerKeysArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.account.hasServerKeys", []interface{}{__arg}, &res)
	return
}

func (c AccountClient) ResetAccount(ctx context.Context, sessionID int) (err error) {
	__arg := ResetAccountArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.account.resetAccount", []interface{}{__arg}, nil)
	return
}
