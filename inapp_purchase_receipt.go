package storekit

// InAppPurchaseReceipt is an array that contains the in-app purchase receipt
// fields for all in-app purchase transactions.
// https://developer.apple.com/documentation/appstorereceipts/responsebody/receipt/in_app
type InAppPurchaseReceipt struct {
	// The time Apple customer support canceled a transaction, or the time an
	// auto-renewable subscription plan was upgraded, in a date-time format similar
	// to the ISO 8601. This field is only present for refunded transactions.
	CancellationDate string `json:"cancellation_date,omitempty"`

	// The time Apple customer support canceled a transaction, or the time an
	// auto-renewable subscription plan was upgraded, in UNIX epoch time format, in
	// milliseconds.This field is only present for refunded transactions.Use this
	// time format for processing dates.
	// information.
	CancellationDateMS int64 `json:"cancellation_date_ms,string,omitempty"`

	// The time Apple customer support canceled a transaction, or the time an
	// auto-renewable subscription plan was upgraded, in the Pacific Time zone. This
	// field is only present for refunded transactions.
	CancellationDatePST string `json:"cancellation_date_pst,omitempty"`

	// The reason for a refunded transaction. When a customer cancels a transaction,
	// the App Store gives them a refund and provides a value for this key. A value
	// of “1” indicates that the customer canceled their transaction due to an
	// actual or perceived issue within your app. A value of “0” indicates that the
	// transaction was canceled for another reason; for example, if the customer
	// made the purchase accidentally.
	CancellationReason int `json:"cancellation_reason,string,omitempty"`

	// The time a subscription expires or when it will renew, in a date-time format
	// similar to the ISO 8601.
	ExpiresDate string `json:"expires_date,omitempty"`

	// The time a subscription expires or when it will renew, in UNIX epoch time
	// format, in milliseconds.Use this time format for processing dates.
	ExpiresDateMS int64 `json:"expires_date_ms,string,omitempty"`

	// The time a subscription expires or when it will renew, in the Pacific Time
	// zone.
	ExpiresDatePST string `json:"expires_date_pst,omitempty"`

	// An indicator of whether an auto-renewable subscription is in the introductory
	// price period.
	IsInIntroOfferPeriod bool `json:"is_in_intro_offer_period,string,omitempty"`

	// An indication of whether a subscription is in the free trial period.
	IsTrialPeriod bool `json:"is_trial_period,string,omitempty"`

	// The time of the original in-app purchase, in a date-time format similar to
	// ISO 8601.
	OriginalPurchaseDate string `json:"original_purchase_date,omitempty"`

	// The time of the original in-app purchase, in UNIX epoch time format, in
	// milliseconds. For an auto-renewable subscription, this value indicates the
	// date of the subscription's initial purchase. The original purchase date
	// applies to all product types and remains the same in all transactions for the
	// same product ID. This value corresponds to the original transaction’s
	// transactionDate property in StoreKit. Use this time format for processing
	// dates.
	OriginalPurchaseDateMS int64 `json:"original_purchase_date_ms,string,omitempty"`

	// The time of the original in-app purchase, in the Pacific Time zone.
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst,omitempty"`

	// OriginalTransactionIdentifier is the transaction identifier of the original
	// transaction for a transaction that restores a previous transaction.
	// Otherwise, identical to the transaction identifier.
	OriginalTransactionID string `json:"original_transaction_id,omitempty"`

	// The unique identifier of the product purchased. You provide this value when
	// creating the product in App Store Connect, and it corresponds to the
	// productIdentifier property of the SKPayment object stored in the
	// transaction's payment property.
	ProductID string `json:"product_id,omitempty"`

	// The identifier of the subscription offer redeemed by the user.
	PromotionalOfferID string `json:"promotional_offer_id,omitempty"`

	// The time the App Store charged the user's account for a purchased or restored
	// product, or the time the App Store charged the user’s account for a
	// subscription purchase or renewal after a lapse, in a date-time format similar
	// to ISO 8601.
	PurchaseDate string `json:"purchase_date,omitempty"`

	// For consumable, non-consumable, and non-renewing subscription products, the
	// time the App Store charged the user's account for a purchased or restored
	// product, in the UNIX epoch time format, in milliseconds. For auto-renewable
	// subscriptions, the time the App Store charged the user’s account for a
	// subscription purchase or renewal after a lapse, in the UNIX epoch time
	// format, in milliseconds. Use this time format for processing dates.
	PurchaseDateMS int64 `json:"purchase_date_ms,string,omitempty"`

	// The time the App Store charged the user's account for a purchased or restored
	// product, or the time the App Store charged the user’s account for a
	// subscription purchase or renewal after a lapse, in the Pacific Time zone.
	PurchaseDatePST string `json:"purchase_date_pst,omitempty"`

	// The number of consumable products purchased. This value corresponds to the
	// quantity property of the SKPayment object stored in the transaction's payment
	// property. The value is usually “1” unless modified with a mutable payment.
	// The maximum value is 10.
	Quantity int `json:"quantity,string,omitempty"`

	// A unique identifier for a transaction such as a purchase, restore, or
	// renewal.
	//
	// This value corresponds to the transaction’s transactionIdentifier property.
	//
	// For a transaction that restores a previous transaction, this value is
	// different from the transaction identifier of the original purchase
	// transaction. In an auto-renewable subscription receipt, a new value for the
	// transaction identifier is generated every time the subscription automatically
	// renews or is restored on a new device.
	TransactionID string `json:"transaction_id,omitempty"`

	// A unique identifier for purchase events across devices, including
	// subscription-renewal events. This value is the primary key for identifying
	// subscription purchases.
	WebOrderLineItemID string `json:"web_order_line_item_id,omitempty"`
}
