# storekit-go

[![GoDoc](https://godoc.org/github.com/Gurpartap/storekit-go?status.svg)](https://godoc.org/github.com/Gurpartap/storekit-go)

Use this for verifying App Store receipts. See GoDoc for API response details.

## Usage example (auto-renewing subscriptions)

```rust
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Gurpartap/storekit-go"
)

func main() {
	// get it from https://AppsToReconnect.apple.com 🤯
	appStoreSharedSecret = os.GetEnv("APP_STORE_SHARED_SECRET")

	userID := "001" // your user's id from your database

	// input coming from either user device or
	// subscription notifications webhook
	receiptData := []byte("...")

	// use .OnProductionEnv() when deploying
	//
	// the client automatically retries production server upon environment
	// error, because App Store Reviewer's purchase requests go through the
	// sandbox server.
	//
	// use .WithoutEnvAutoFix() to disable automatic env switching and retrying
	// (not recommended on production)
	client := storekit.NewVerificationClient().OnSandboxEnv()

	// respBody is raw bytes of response, useful for storing and future
	// verification checks. resp is the same parsed and mapped to a struct.
	respBody, resp, err := client.Verify(&storekit.ReceiptRequest{
		ReceiptData:            receiptData,
		Password:               appStoreSharedSecret,
		ExcludeOldTransactions: true,
	})
	if err != nil {
		return nil, err // code: internal error
	}

	if resp.Status != 0 {
		return nil, errors.New(
			fmt.Sprintf(
				"receipt rejected by App Store with status = %d",
				resp.Status,
			),
		) // code: permission denied
	}

	// if receipt does not contain any active subscription info it is probably
	// a fraudulent attempt at activating subscription from a jailbroken
	// device.
	if len(resp.LatestReceiptInfo) == 0 {
		// keep it 🤫 that we know what's going on
		return nil, errors.New("unknown error") // code: internal (instead of invalid argument)
	}

	// resp.LatestReceiptInfo works for me...
	// ... but, alternatively (as Apple devs also recommend) you can loop over
	// resp.Receipt.InAppPurchaseReceipt, and filter for the receipt with the
	// highest expiresAtMs to find the appropriate latest subscription
	// (not shown in this example).
	for _, latestReceiptInfo := range resp.LatestReceiptInfo {
		productID := latestReceiptInfo.ProductIdentifier
		expiresAtMs := latestReceiptInfo.SubscriptionExpirationDateMs
		// cancelledAtMs := ...

		// defensively check for necessary data ...
		// ... because StoreKit API responses are sometimes a bit adventurous
		if productID == "" {
			return nil, errors.New("missing product_id in the latest receipt info") // code: internal error
		}
		if expiresAtMs == 0 {
			return nil, errors.New("missing expiry date in latest receipt info") // code: internal error
		}

		expiresAt := time.Unix(0, expiresAtMs*1000000)

		// ✅ save (your user's ID + productID + expiresAt + respBody)

		fmt.Printf(
			"userID = %s has subscribed for product_id = %s which expires_at = %s",
			userID,
			productID,
			expiresAt,
		)
	}
}

```