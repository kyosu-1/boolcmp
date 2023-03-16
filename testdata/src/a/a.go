package a

func f() {
	isOpen := true
	if isOpen == true { // want "bool value used in comparison"
		// do something
	}
	
	isClosed := false
	if isClosed == false { // want "bool value used in comparison"
		// do something else
	}
	
	isEnabled := true
	if isEnabled { // ok
		// do something else
	}
	
	isDisabled := false
	if !isDisabled { // ok
		// do something else
	}

	if isOpen == isClosed { // want "bool value used in comparison"
		// do something else
	}

	if true == isOpen { // want "bool value used in comparison"
		// do something else
	}
}
