// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package namegenerator

import (
	_ "embed"
	"math/rand"
	"strings"
)

var (
	//go:embed wordlists/adjectives.txt
	adj string
	//go:embed wordlists/nouns.txt
	noun string
)

// Generator ...
type Generator interface {
	Generate(int, string) string
}

// NameGenerator ...
type NameGenerator struct {
	random *rand.Rand
	adj    []string
	noun   []string
}

// Generate ...
func (rn *NameGenerator) Generate(count int, divider string) string {

	randomName := ""
	for i := 0; i < count-1; i++ {
		randomName += rn.adj[rn.random.Intn(len(rn.adj))] + divider
	}
	randomName += rn.noun[rn.random.Intn(len(rn.noun))]
	
	return randomName
}

// NewNameGenerator ...
func NewNameGenerator(seed int64) Generator {
	nameGenerator := &NameGenerator{
		random: rand.New(rand.New(rand.NewSource(99))),
		adj:    strings.Fields(adj),
		noun:   strings.Fields(noun),
	}
	nameGenerator.random.Seed(seed)

	return nameGenerator
}
