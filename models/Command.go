package models

type Command struct {
	CType CommandType

	/**
	 *  request unique identification
	 */
	opaque uint64

	/**
	 *  data body
	 */
	body []byte
}

func (c *Command) GetBody() []byte {
	return c.body
}
