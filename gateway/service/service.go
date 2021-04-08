package service

import (
	"encoding/json"
	"gateway/service/loadbalance"
	"gateway/upstream"
)

type Service struct {
	LoadBalance loadbalance.LoadBalance
	HashOn      string
	Nodes       []upstream.Node
}

const (
	RANDOM      = "random"
	RoundRibbon = "round_ribbon"
	HASH        = "hash"
)

func (s *Service) UnmarshalJSON(data []byte) error {
	ss := struct {
		LoadBalanceType string `json:"load_balance_type"`
		HashOn          string
		Nodes           []upstream.Node `json:"nodes"`
	}{}

	if err := json.Unmarshal(data, &ss); err != nil {
		return err
	}
	s.Nodes = ss.Nodes
	switch ss.LoadBalanceType {
	case RANDOM:
		s.LoadBalance = &loadbalance.RandomLoadBalance{}
	case RoundRibbon:
		s.LoadBalance = &loadbalance.RoundRibbonLoadBalance{}
	case HASH:
		s.LoadBalance = &loadbalance.HashLoadBalance{
			HashOn: ss.HashOn,
		}
	}
	return nil
}

func (s Service) MarshalJSON() ([]byte, error) {
	var lbtype = ""
	var hashOn = ""
	switch s.LoadBalance.(type) {
	case *loadbalance.RandomLoadBalance:
		lbtype = RANDOM
	case *loadbalance.RoundRibbonLoadBalance:
		lbtype = RoundRibbon
	case *loadbalance.HashLoadBalance:
		lbtype = HASH
		hashOn = s.LoadBalance.(*loadbalance.HashLoadBalance).HashOn
	}
	return json.Marshal(struct {
		LoadBalanceType string `json:"load_balance_type"`
		HashOn          string
		Nodes           []upstream.Node `json:"nodes"`
	}{
		LoadBalanceType: lbtype,
		HashOn:          hashOn,
		Nodes:           s.Nodes,
	})
}
