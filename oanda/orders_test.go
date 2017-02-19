// Copyright 2014 Tjerk Santegoeds
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oanda_test

import (
	"time"

	"github.com/santegoeds/oanda"

	"gopkg.in/check.v1"
)

type TestOrderSuite struct {
	OandaSuite
}

func (ts *TestOrderSuite) SetUpSuite(c *check.C) {
	ts.OandaSuite.SetUpSuite(c)
	ts.SetUpAccount(c)
}

var _ = check.Suite(&TestOrderSuite{})

func (ts *TestOrderSuite) TestOrderApi(c *check.C) {
	expiry := time.Now().Add(5 * time.Minute)

	o, err := ts.Client.NewOrder(oanda.Limit, oanda.Buy, 2, "eur_usd", 0.75, expiry,
		oanda.UpperBound(1.0), oanda.LowerBound(0.5))
	c.Assert(err, check.IsNil)
	c.Log(o)
	c.Assert(o.OrderId, check.Not(check.Equals), 0)
	orderExpiry := o.Expiry.Time()
	c.Assert(orderExpiry.Equal(expiry.Truncate(time.Second)), check.Equals, true)
	c.Assert(o.Instrument, check.Equals, "EUR_USD")
	c.Assert(o.OrderType, check.Equals, string(oanda.Limit))
	c.Assert(o.Price, check.Equals, 0.75)
	c.Assert(o.Units, check.Equals, 2)
	c.Assert(o.Side, check.Equals, string(oanda.Buy))
	c.Assert(o.LowerBound, check.Equals, 0.5)
	c.Assert(o.UpperBound, check.Equals, 1.0)
	c.Assert(o.StopLoss, check.Equals, 0.0)
	c.Assert(o.TakeProfit, check.Equals, 0.0)
	c.Assert(o.TrailingStop, check.Equals, 0.0)

	dup, err := ts.Client.Order(o.OrderId)
	c.Assert(err, check.IsNil)
	c.Assert(dup.OrderId, check.Equals, o.OrderId)
	c.Assert(dup.Expiry, check.Equals, o.Expiry)
	c.Assert(dup.Instrument, check.Equals, o.Instrument)
	c.Assert(dup.OrderType, check.Equals, o.OrderType)
	c.Assert(dup.Price, check.Equals, o.Price)
	c.Assert(dup.Units, check.Equals, o.Units)
	c.Assert(dup.Side, check.Equals, o.Side)
	c.Assert(dup.LowerBound, check.Equals, o.LowerBound)
	c.Assert(dup.UpperBound, check.Equals, o.UpperBound)
	c.Assert(dup.StopLoss, check.Equals, o.StopLoss)
	c.Assert(dup.TakeProfit, check.Equals, o.TakeProfit)
	c.Assert(dup.TrailingStop, check.Equals, o.TrailingStop)

	orders, err := ts.Client.Orders()
	c.Assert(err, check.IsNil)
	c.Log(orders)
	c.Assert(orders, check.HasLen, 1)
	c.Assert(orders[0].OrderId, check.Equals, o.OrderId)
	c.Assert(orders[0].Expiry, check.Equals, o.Expiry)
	c.Assert(orders[0].Instrument, check.Equals, o.Instrument)
	c.Assert(orders[0].OrderType, check.Equals, o.OrderType)
	c.Assert(orders[0].Price, check.Equals, o.Price)
	c.Assert(orders[0].Units, check.Equals, o.Units)
	c.Assert(orders[0].Side, check.Equals, o.Side)
	c.Assert(orders[0].LowerBound, check.Equals, o.LowerBound)
	c.Assert(orders[0].UpperBound, check.Equals, o.UpperBound)
	c.Assert(orders[0].StopLoss, check.Equals, o.StopLoss)
	c.Assert(orders[0].TakeProfit, check.Equals, o.TakeProfit)
	c.Assert(orders[0].TrailingStop, check.Equals, o.TrailingStop)

	o, err = ts.Client.ModifyOrder(o.OrderId, oanda.Units(1))
	c.Assert(err, check.IsNil)
	c.Assert(o.Units, check.Equals, 1)

	rsp, err := ts.Client.CancelOrder(o.OrderId)
	c.Assert(err, check.IsNil)
	c.Log("OrderCancelResponse:", rsp)
	orders, err = ts.Client.Orders()
	c.Assert(err, check.IsNil)
	c.Assert(orders, check.HasLen, 0)
}
