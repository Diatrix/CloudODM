// Copyright © 2018 CloudODM Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package odm

import (
	"testing"
)

func TestAPI(t *testing.T) {
	node := Node{URL: "http://localhost:3000"}
	offlineNode := Node{URL: "http://unknownhost:3000"}

	resp, err := node.Info()
	if err != nil {
		t.Error("Cannot retrieve /info")
	}

	if resp.Version == "" {
		t.Error("Version is not set")
	}

	resp, err = offlineNode.Info()
	if err == nil {
		t.Error("No error when retrieving an offline node /info")
	}

	node._debugUnauthorized = true
	resp, err = node.Info()
	t.Log(resp, err)
	if err != ErrUnauthorized {
		t.Error("Error should have been ErrUnauthorized")
	}
}
