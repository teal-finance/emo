// Code generated by https://github.com/teal-finance/emo/blob/main/codegen/golang/gen.go ; DO NOT EDIT.

package emo

func (zone Zone) Info(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("ℹ️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Warning(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔔", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Error(args ...any) Event {
	return new("💢", zone, true, args).print().callHook()
}

func (zone Zone) Query(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🗄️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) QueryError(args ...any) Event {
	return new("🗄️", zone, true, args).print().callHook()
}

func (zone Zone) Encrypt(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🎼", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) EncryptError(args ...any) Event {
	return new("🎼", zone, true, args).print().callHook()
}

func (zone Zone) Decrypt(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🗝️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) DecryptError(args ...any) Event {
	return new("🗝️", zone, true, args).print().callHook()
}

func (zone Zone) Time(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("⏱️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) TimeError(args ...any) Event {
	return new("⏱️", zone, true, args).print().callHook()
}

func (zone Zone) Param(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📩", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) ParamError(args ...any) Event {
	return new("📩", zone, true, args).print().callHook()
}

func (zone Zone) Debug(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("💊", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) State(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📢", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Save(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("💾", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Delete(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("❌", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Data(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("💼", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Line(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("➖", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Init(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🎬", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Update(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🆙", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Ok(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🆗", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Build(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔧", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Aconstructor(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🛠️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) NotFound(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🚫", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Found(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("👁️‍🗨️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Result(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📌", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Input(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📥", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Output(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📤", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Function(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔨", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Key(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔑", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) AccessToken(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔑", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) RefreshToken(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🗝️", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Transmit(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("📡", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Start(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🏁", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) Stop(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🛑", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) ArrowIn(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("=>", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) ArrowOut(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("<=", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) SmallArrowIn(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("->", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) SmallArrowOut(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("<-", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) RequestGet(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔷", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}

func (zone Zone) RequestPost(args ...any) Event {
	if zone.Print || (zone.Hook != nil) {
		return new("🔶", zone, false, args).print().callHook()
	}
	var evt Event
	return evt
}
