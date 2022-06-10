package element

// Coder serializes the input Element into binary
type Coder func(Element) ([]byte, error)

// Decoder deserializes the input binary into Element.
// If the input Element is not nil, Decoder should check and use it, but not create a new one.
type Decoder func([]byte, Element) (Element, error)
