package sexpconv

import (
	"lisp/function"
	"sexp"
)

// Simplify translates semantic-rich forms into
// more generic and fundamental forms.
// For example, it converts DoTimes to While.
//
// It is necessary to call Simplify before compiling
// forms into IR because compiler does not recognize
// some high level constructs.
func Simplify(form sexp.Form) sexp.Form {
	return sexp.Rewrite(form, simplify)
}

func simplifyList(forms []sexp.Form) []sexp.Form {
	for i, form := range forms {
		forms[i] = sexp.Rewrite(form, simplify)
	}
	return forms
}

func simplify(form sexp.Form) sexp.Form {
	switch form := form.(type) {
	case *sexp.SliceLit:
		return &sexp.LispCall{
			Fn: function.ArrayToSlice,
			Args: []sexp.Form{
				sexp.NewLispCall(function.Vector, simplifyList(form.Vals)...),
			},
		}

	case *sexp.ArrayLit:
		return sexp.NewLispCall(function.Vector, simplifyList(form.Vals)...)

	case *sexp.DoTimes:
		bindKey := &sexp.Bind{
			Name: form.Iter.Name,
			Init: ZeroValue(form.Iter.Typ),
		}
		form.Body.Forms = append(form.Body.Forms, &sexp.Rebind{
			Name: form.Iter.Name,
			Expr: sexp.NewAdd1(form.Iter),
		})
		loop := &sexp.While{
			Cond: sexp.NewNumLt(form.Iter, form.N),
			Body: form.Body,
		}
		return &sexp.Block{
			Forms: []sexp.Form{bindKey, loop},
		}
	}

	return nil
}
