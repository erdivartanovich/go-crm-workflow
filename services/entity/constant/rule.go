package constant

const (
	/**
	 * Field Types.
	 */
	FIELD_TYPE_BOOLEAN  = 1
	FIELD_TYPE_CURRENCY = 2
	FIELD_TYPE_DATE     = 3
	FIELD_TYPE_NUMBER   = 4
	FIELD_TYPE_PERCENT  = 5
	FIELD_TYPE_MONTH    = 6

	/**
	 * Operators.
	 */
	OPERATOR_EQUAL                 = 1  // =
	OPERATOR_NOT_EQUAL             = 2  // <> | !=
	OPERATOR_GREATER_THAN_OR_EQUAL = 3  // >=
	OPERATOR_GREATER_THAN          = 4  // >
	OPERATOR_LESS_THAN_OR_EQUAL    = 5  // <=
	OPERATOR_LESS_THAN             = 6  // <
	OPERATOR_BETWEEN               = 7  // between
	OPERATOR_NOT_BETWEEN           = 8  // not between
	OPERATOR_IN                    = 9  // in()
	OPERATOR_NOT_IN                = 10 //not_in

	/**
	 * Runnable interval.
	 */
	RUNNABLE_INTERVAL_DAILY   = 1
	RUNNABLE_INTERVAL_MONTHLY = 2
	RUNNABLE_INTERVAL_ANNUAL  = 3
)
