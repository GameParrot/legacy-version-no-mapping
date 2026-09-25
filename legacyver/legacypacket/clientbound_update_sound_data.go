package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientboundUpdateSoundData(io protocol.IO, pk *packet.ClientboundUpdateSoundData) {
	io.Uint64(&pk.ServerSoundHandle)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.Single(io, &pk.Stop)
		protocol.Single(io, &pk.SetVolume)
		protocol.Single(io, &pk.SetPitch)
		protocol.Single(io, &pk.Fade)
		protocol.Single(io, &pk.SeekTo)
		protocol.Single(io, &pk.Pause)
		protocol.Single(io, &pk.Resume)
		return
	}
	event := "Stop"
	io.String(&event)
	if proto.IsReader(io) {
		if event != "Stop" {
			io.UnknownEnumOption(event, "sound data event")
			return
		}
		pk.Stop = protocol.SoundDataUpdate{Type: protocol.SoundDataUpdateStop}
	}
}
