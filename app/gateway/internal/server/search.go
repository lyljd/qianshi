package server

import "qianshi/app/gateway/internal/router"

func search(routers []router.Router, paths []string) (find bool, to string, needLogin bool, power int) {
	if routers == nil || len(paths) == 0 {
		return
	}

	path := paths[0]
	for _, r := range routers {
		if r.Path == path {
			_, cTo, cAuth, cPower := search(r.Children, paths[1:])
			find, needLogin, power = true, cAuth || r.NeedLogin, max(cPower, r.Power)
			if cTo != "" {
				to = cTo
			} else {
				to = r.To
			}
			break
		}
	}

	return
}
