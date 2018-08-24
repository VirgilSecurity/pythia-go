/*
 * Copyright (C) 2015-2018 Virgil Security Inc.
 *
 * Lead Maintainer: Virgil Security Inc. <support@virgilsecurity.com>
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *   (1) Redistributions of source code must retain the above copyright
 *   notice, this list of conditions and the following disclaimer.
 *
 *   (2) Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in
 *   the documentation and/or other materials provided with the
 *   distribution.
 *
 *   (3) Neither the name of the copyright holder nor the names of its
 *   contributors may be used to endorse or promote products derived
 *   from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ''AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

package sdk

import "gopkg.in/virgil.v4/errors"

const (
	VirgilContentType = "virgil-jwt;v=1"
	JwtType           = "JWT"
)

type JwtHeaderContent struct {
	Algorithm   string `json:"alg"`
	Type        string `json:"typ"`
	ContentType string `json:"cty"`
	ApiKeyId    string `json:"kid"`
}

func NewJwtHeaderContent(algorithm string, apiKeyId string) (*JwtHeaderContent, error) {
	if SpaceMap(algorithm) == "" {
		return nil, errors.New("algorithm is empty")
	}
	if SpaceMap(apiKeyId) == "" {
		return nil, errors.New("apiKeyId is empty")
	}
	return &JwtHeaderContent{
		Algorithm:   algorithm,
		ApiKeyId:    apiKeyId,
		ContentType: VirgilContentType,
		Type:        JwtType,
	}, nil
}
