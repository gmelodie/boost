// Code generated by github.com/filecoin-project/lotus/gen/api. DO NOT EDIT.

package api

import (
	"context"

	"github.com/filecoin-project/boost/journal/alerting"
	"github.com/filecoin-project/go-jsonrpc/auth"
	apitypes "github.com/filecoin-project/lotus/api/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	xerrors "golang.org/x/xerrors"
)

var ErrNotSupported = xerrors.New("method not supported")

type BoostStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
	}
}

type BoostStub struct {
	CommonStub

	NetStub
}

type ChainIOStruct struct {
	Internal struct {
		ChainHasObj func(p0 context.Context, p1 cid.Cid) (bool, error) ``

		ChainReadObj func(p0 context.Context, p1 cid.Cid) ([]byte, error) ``
	}
}

type ChainIOStub struct {
}

type CommonStruct struct {
	Internal struct {
		AuthNew func(p0 context.Context, p1 []auth.Permission) ([]byte, error) `perm:"admin"`

		AuthVerify func(p0 context.Context, p1 string) ([]auth.Permission, error) `perm:"read"`

		Closing func(p0 context.Context) (<-chan struct{}, error) `perm:"read"`

		Discover func(p0 context.Context) (apitypes.OpenRPCDocument, error) `perm:"read"`

		LogAlerts func(p0 context.Context) ([]alerting.Alert, error) `perm:"admin"`

		LogList func(p0 context.Context) ([]string, error) `perm:"write"`

		LogSetLevel func(p0 context.Context, p1 string, p2 string) error `perm:"write"`

		Session func(p0 context.Context) (uuid.UUID, error) `perm:"read"`

		Shutdown func(p0 context.Context) error `perm:"admin"`

		Version func(p0 context.Context) (APIVersion, error) `perm:"read"`
	}
}

type CommonStub struct {
}

type CommonNetStruct struct {
	CommonStruct

	NetStruct

	Internal struct {
	}
}

type CommonNetStub struct {
	CommonStub

	NetStub
}

type NetStruct struct {
	Internal struct {
		ID func(p0 context.Context) (peer.ID, error) `perm:"read"`

		NetAddrsListen func(p0 context.Context) (peer.AddrInfo, error) `perm:"read"`

		NetAgentVersion func(p0 context.Context, p1 peer.ID) (string, error) `perm:"read"`

		NetAutoNatStatus func(p0 context.Context) (NatInfo, error) `perm:"read"`

		NetBandwidthStats func(p0 context.Context) (metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByPeer func(p0 context.Context) (map[string]metrics.Stats, error) `perm:"read"`

		NetBandwidthStatsByProtocol func(p0 context.Context) (map[protocol.ID]metrics.Stats, error) `perm:"read"`

		NetBlockAdd func(p0 context.Context, p1 NetBlockList) error `perm:"admin"`

		NetBlockList func(p0 context.Context) (NetBlockList, error) `perm:"read"`

		NetBlockRemove func(p0 context.Context, p1 NetBlockList) error `perm:"admin"`

		NetConnect func(p0 context.Context, p1 peer.AddrInfo) error `perm:"write"`

		NetConnectedness func(p0 context.Context, p1 peer.ID) (network.Connectedness, error) `perm:"read"`

		NetDisconnect func(p0 context.Context, p1 peer.ID) error `perm:"write"`

		NetFindPeer func(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) `perm:"read"`

		NetPeerInfo func(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) `perm:"read"`

		NetPeers func(p0 context.Context) ([]peer.AddrInfo, error) `perm:"read"`

		NetPubsubScores func(p0 context.Context) ([]PubsubScore, error) `perm:"read"`
	}
}

type NetStub struct {
}

func (s *ChainIOStruct) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	if s.Internal.ChainHasObj == nil {
		return false, ErrNotSupported
	}
	return s.Internal.ChainHasObj(p0, p1)
}

func (s *ChainIOStub) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, ErrNotSupported
}

func (s *ChainIOStruct) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	if s.Internal.ChainReadObj == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.ChainReadObj(p0, p1)
}

func (s *ChainIOStub) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	if s.Internal.AuthNew == nil {
		return *new([]byte), ErrNotSupported
	}
	return s.Internal.AuthNew(p0, p1)
}

func (s *CommonStub) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	return *new([]byte), ErrNotSupported
}

func (s *CommonStruct) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	if s.Internal.AuthVerify == nil {
		return *new([]auth.Permission), ErrNotSupported
	}
	return s.Internal.AuthVerify(p0, p1)
}

func (s *CommonStub) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	return *new([]auth.Permission), ErrNotSupported
}

func (s *CommonStruct) Closing(p0 context.Context) (<-chan struct{}, error) {
	if s.Internal.Closing == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.Closing(p0)
}

func (s *CommonStub) Closing(p0 context.Context) (<-chan struct{}, error) {
	return nil, ErrNotSupported
}

func (s *CommonStruct) Discover(p0 context.Context) (apitypes.OpenRPCDocument, error) {
	if s.Internal.Discover == nil {
		return *new(apitypes.OpenRPCDocument), ErrNotSupported
	}
	return s.Internal.Discover(p0)
}

func (s *CommonStub) Discover(p0 context.Context) (apitypes.OpenRPCDocument, error) {
	return *new(apitypes.OpenRPCDocument), ErrNotSupported
}

func (s *CommonStruct) LogAlerts(p0 context.Context) ([]alerting.Alert, error) {
	if s.Internal.LogAlerts == nil {
		return *new([]alerting.Alert), ErrNotSupported
	}
	return s.Internal.LogAlerts(p0)
}

func (s *CommonStub) LogAlerts(p0 context.Context) ([]alerting.Alert, error) {
	return *new([]alerting.Alert), ErrNotSupported
}

func (s *CommonStruct) LogList(p0 context.Context) ([]string, error) {
	if s.Internal.LogList == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.LogList(p0)
}

func (s *CommonStub) LogList(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *CommonStruct) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	if s.Internal.LogSetLevel == nil {
		return ErrNotSupported
	}
	return s.Internal.LogSetLevel(p0, p1, p2)
}

func (s *CommonStub) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	return ErrNotSupported
}

func (s *CommonStruct) Session(p0 context.Context) (uuid.UUID, error) {
	if s.Internal.Session == nil {
		return *new(uuid.UUID), ErrNotSupported
	}
	return s.Internal.Session(p0)
}

func (s *CommonStub) Session(p0 context.Context) (uuid.UUID, error) {
	return *new(uuid.UUID), ErrNotSupported
}

func (s *CommonStruct) Shutdown(p0 context.Context) error {
	if s.Internal.Shutdown == nil {
		return ErrNotSupported
	}
	return s.Internal.Shutdown(p0)
}

func (s *CommonStub) Shutdown(p0 context.Context) error {
	return ErrNotSupported
}

func (s *CommonStruct) Version(p0 context.Context) (APIVersion, error) {
	if s.Internal.Version == nil {
		return *new(APIVersion), ErrNotSupported
	}
	return s.Internal.Version(p0)
}

func (s *CommonStub) Version(p0 context.Context) (APIVersion, error) {
	return *new(APIVersion), ErrNotSupported
}

func (s *NetStruct) ID(p0 context.Context) (peer.ID, error) {
	if s.Internal.ID == nil {
		return *new(peer.ID), ErrNotSupported
	}
	return s.Internal.ID(p0)
}

func (s *NetStub) ID(p0 context.Context) (peer.ID, error) {
	return *new(peer.ID), ErrNotSupported
}

func (s *NetStruct) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	if s.Internal.NetAddrsListen == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetAddrsListen(p0)
}

func (s *NetStub) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	if s.Internal.NetAgentVersion == nil {
		return "", ErrNotSupported
	}
	return s.Internal.NetAgentVersion(p0, p1)
}

func (s *NetStub) NetAgentVersion(p0 context.Context, p1 peer.ID) (string, error) {
	return "", ErrNotSupported
}

func (s *NetStruct) NetAutoNatStatus(p0 context.Context) (NatInfo, error) {
	if s.Internal.NetAutoNatStatus == nil {
		return *new(NatInfo), ErrNotSupported
	}
	return s.Internal.NetAutoNatStatus(p0)
}

func (s *NetStub) NetAutoNatStatus(p0 context.Context) (NatInfo, error) {
	return *new(NatInfo), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	if s.Internal.NetBandwidthStats == nil {
		return *new(metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStats(p0)
}

func (s *NetStub) NetBandwidthStats(p0 context.Context) (metrics.Stats, error) {
	return *new(metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByPeer == nil {
		return *new(map[string]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByPeer(p0)
}

func (s *NetStub) NetBandwidthStatsByPeer(p0 context.Context) (map[string]metrics.Stats, error) {
	return *new(map[string]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	if s.Internal.NetBandwidthStatsByProtocol == nil {
		return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
	}
	return s.Internal.NetBandwidthStatsByProtocol(p0)
}

func (s *NetStub) NetBandwidthStatsByProtocol(p0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	return *new(map[protocol.ID]metrics.Stats), ErrNotSupported
}

func (s *NetStruct) NetBlockAdd(p0 context.Context, p1 NetBlockList) error {
	if s.Internal.NetBlockAdd == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockAdd(p0, p1)
}

func (s *NetStub) NetBlockAdd(p0 context.Context, p1 NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetBlockList(p0 context.Context) (NetBlockList, error) {
	if s.Internal.NetBlockList == nil {
		return *new(NetBlockList), ErrNotSupported
	}
	return s.Internal.NetBlockList(p0)
}

func (s *NetStub) NetBlockList(p0 context.Context) (NetBlockList, error) {
	return *new(NetBlockList), ErrNotSupported
}

func (s *NetStruct) NetBlockRemove(p0 context.Context, p1 NetBlockList) error {
	if s.Internal.NetBlockRemove == nil {
		return ErrNotSupported
	}
	return s.Internal.NetBlockRemove(p0, p1)
}

func (s *NetStub) NetBlockRemove(p0 context.Context, p1 NetBlockList) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	if s.Internal.NetConnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetConnect(p0, p1)
}

func (s *NetStub) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	return ErrNotSupported
}

func (s *NetStruct) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	if s.Internal.NetConnectedness == nil {
		return *new(network.Connectedness), ErrNotSupported
	}
	return s.Internal.NetConnectedness(p0, p1)
}

func (s *NetStub) NetConnectedness(p0 context.Context, p1 peer.ID) (network.Connectedness, error) {
	return *new(network.Connectedness), ErrNotSupported
}

func (s *NetStruct) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	if s.Internal.NetDisconnect == nil {
		return ErrNotSupported
	}
	return s.Internal.NetDisconnect(p0, p1)
}

func (s *NetStub) NetDisconnect(p0 context.Context, p1 peer.ID) error {
	return ErrNotSupported
}

func (s *NetStruct) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	if s.Internal.NetFindPeer == nil {
		return *new(peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetFindPeer(p0, p1)
}

func (s *NetStub) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	return *new(peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetPeerInfo(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) {
	if s.Internal.NetPeerInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.NetPeerInfo(p0, p1)
}

func (s *NetStub) NetPeerInfo(p0 context.Context, p1 peer.ID) (*ExtendedPeerInfo, error) {
	return nil, ErrNotSupported
}

func (s *NetStruct) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	if s.Internal.NetPeers == nil {
		return *new([]peer.AddrInfo), ErrNotSupported
	}
	return s.Internal.NetPeers(p0)
}

func (s *NetStub) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	return *new([]peer.AddrInfo), ErrNotSupported
}

func (s *NetStruct) NetPubsubScores(p0 context.Context) ([]PubsubScore, error) {
	if s.Internal.NetPubsubScores == nil {
		return *new([]PubsubScore), ErrNotSupported
	}
	return s.Internal.NetPubsubScores(p0)
}

func (s *NetStub) NetPubsubScores(p0 context.Context) ([]PubsubScore, error) {
	return *new([]PubsubScore), ErrNotSupported
}

var _ Boost = new(BoostStruct)
var _ ChainIO = new(ChainIOStruct)
var _ Common = new(CommonStruct)
var _ CommonNet = new(CommonNetStruct)
var _ Net = new(NetStruct)
