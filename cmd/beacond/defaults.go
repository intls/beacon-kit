// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package main

import (
	"github.com/berachain/beacon-kit/node-core/components"
)

func DefaultComponents() []any {
	c := []any{
		components.ProvideAttributesFactory[*Logger],
		components.ProvideAvailabilityStore[*Logger],
		components.ProvideDepositContract,
		components.ProvideBlockStore[*Logger],
		components.ProvideBlsSigner,
		components.ProvideBlobProcessor[*ConsensusSidecars, *Logger],
		components.ProvideBlobProofVerifier,
		components.ProvideChainService[
			*ConsensusBlock,
			*ConsensusSidecars,
			*DepositContract,
			*Genesis,
			*KVStore, *Logger, *StorageBackend,
		],
		components.ProvideNode,
		components.ProvideChainSpec,
		components.ProvideConfig,
		components.ProvideServerConfig,
		// components.ProvideConsensusEngine[
		// 	*AvailabilityStore, *BeaconBlockHeader, *BeaconState,
		// 	*BeaconStateMarshallable, *BlockStore, *KVStore, *StorageBackend,
		// ],
		components.ProvideDepositStore[*Logger],
		components.ProvideEngineClient[*Logger],
		components.ProvideExecutionEngine[*Logger],
		components.ProvideJWTSecret,
		components.ProvideLocalBuilder[
			*KVStore, *Logger,
		],
		components.ProvideReportingService[*Logger],
		components.ProvideCometBFTService[*Logger],
		components.ProvideServiceRegistry[
			*ConsensusBlock,
			*ConsensusSidecars,
			*Genesis, *KVStore, *Logger,
			NodeAPIContext,
		],
		components.ProvideSidecarFactory,
		components.ProvideStateProcessor[
			*Logger,
			*KVStore,
		],
		components.ProvideKVStore,
		components.ProvideStorageBackend[*KVStore],
		components.ProvideTelemetrySink,
		components.ProvideTelemetryService,
		components.ProvideTrustedSetup,
		components.ProvideValidatorService[
			*KVStore, *Logger,
			*StorageBackend,
		],
		// TODO Hacks
		components.ProvideKVStoreService,
		components.ProvideKVStoreKey,
	}
	c = append(c,
		components.ProvideNodeAPIServer[*Logger, NodeAPIContext],
		components.ProvideNodeAPIEngine,
		components.ProvideNodeAPIBackend[
			*KVStore, *CometBFTService, *StorageBackend,
		],
	)

	c = append(c,
		components.ProvideNodeAPIHandlers[
			*KVStore, NodeAPIContext,
		],
		components.ProvideNodeAPIBeaconHandler[
			*CometBFTService, NodeAPIContext,
		],
		components.ProvideNodeAPIBuilderHandler[NodeAPIContext],
		components.ProvideNodeAPIConfigHandler[NodeAPIContext],
		components.ProvideNodeAPIDebugHandler[NodeAPIContext],
		components.ProvideNodeAPIEventsHandler[NodeAPIContext],
		components.ProvideNodeAPINodeHandler[NodeAPIContext],
		components.ProvideNodeAPIProofHandler[
			*KVStore, *CometBFTService, NodeAPIContext,
		],
	)

	return c
}
