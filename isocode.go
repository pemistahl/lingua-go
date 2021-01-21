/*
 * Copyright © 2021 Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua

//go:generate stringer -type=IsoCode639_1
type IsoCode639_1 int

//go:generate stringer -type=IsoCode639_3
type IsoCode639_3 int

const (
	AF IsoCode639_1 = iota
	AR
	AZ
	BE
	BG
	BN
	BS
	CA
	CS
	CY
	DA
	DE
	EL
	EN
	EO
	ES
	ET
	EU
	FA
	FI
	FR
	GA
	GU
	HE
	HI
	HR
	HU
	HY
	ID
	IS
	IT
	JA
	KA
	KK
	KO
	LA
	LG
	LT
	LV
	MK
	MN
	MR
	MS
	NB
	NL
	NN
	PA
	PL
	PT
	RO
	RU
	SK
	SL
	SN
	SO
	SQ
	SR
	ST
	SV
	SW
	TA
	TE
	TH
	TL
	TN
	TR
	TS
	UK
	UR
	VI
	XH
	YO
	ZH
	ZU
)

const (
	AFR IsoCode639_3 = iota
	ARA
	AZE
	BEL
	BEN
	BOS
	BUL
	CAT
	CES
	CYM
	DAN
	DEU
	ELL
	ENG
	EPO
	EST
	EUS
	FAS
	FIN
	FRA
	GLE
	GUJ
	HEB
	HIN
	HRV
	HUN
	HYE
	IND
	ISL
	ITA
	JPN
	KAT
	KAZ
	KOR
	LAT
	LAV
	LIT
	LUG
	MAR
	MKD
	MON
	MSA
	NLD
	NNO
	NOB
	PAN
	POL
	POR
	RON
	RUS
	SLK
	SLV
	SNA
	SOM
	SOT
	SPA
	SQI
	SRP
	SWA
	SWE
	TAM
	TEL
	TGL
	THA
	TSN
	TSO
	TUR
	UKR
	URD
	VIE
	XHO
	YOR
	ZHO
	ZUL
)
