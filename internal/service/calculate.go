package service

import (
	"fmt"

	"github.com/sluga1naroda/template-web-server/internal/model"
)

//1012: Даны вещественные числа: A, B, C. Определить, выполняются ли неравенства A < B < C или A < B > C и какое именно неравенство выполняется.

func Calculate(request *model.RequestTask) model.ResponseCalculation {
	var expression string

	if request.A < request.B && request.B < request.C {
		expression = fmt.Sprintf("%f < %f < %f", request.A, request.B, request.C)
	} else if request.A < request.B && request.B > request.C {
		expression = fmt.Sprintf("%f < %f > %f", request.A, request.B, request.C)
	} else if request.A > request.B && request.B > request.C {
		expression = fmt.Sprintf("%f > %f > %f", request.A, request.B, request.C)
	} else if request.A > request.B && request.B < request.C {
		expression = fmt.Sprintf("%f > %f < %f", request.A, request.B, request.C)
	}

	return model.ResponseCalculation{
		Result: expression,
	}
}
