## Note To Self 

-Test Driven Development: Write unit tests before implementing anything  

## ToDo 

1. Analyze the project 
2. Design the architecture ## Map out the parts of a typical bc project and go from there
3. Implement based on existing work. 


## Analysis 

This project is going to base it's design on the bitcoin core protocol

What's a blockchain: a distributed ledger   
That stores all the transactions that ever happened among participants 
since it's inception.

The distributed ledger has a native currency that allow it's users to exchange 
value


### Hashes 

A hash is a string representation of data 

## Features of our Blockchain 
    - A blockchain node 
    - P2P capabilities 
    - Native Currency 
    - Consensu algorithm (PoW)



# System Architecture 

## Components 

    - The blockchain server 
    - The wallet server
    - Mining algo 
    - Consensus algo 
    - Crypto modules for cryptographic algorithms



## Ethereum and Bitcoin Module Benchmark 

### Ethereum 
    - **account**: high level Ethereum account management 

    - **cmd**: command line utilities for the eth node

    - **common**: this package implements common utilities used accross the codebase (Utilities)

    - **consensus**: this package implements different Ethereum consensus engines

    -**console**: A javascript interpreted runtime environment. It is a fully 
    fledged javascript console attached to a running node via na external or 
    in-process RPC 

    - **contracts/checkpointoracle**: An onchain light client checkpoint oracle
    
    - **Core**: the core implementation of the blockchain
    
    - **Crypto**: cryptographic algorithms used in the project 
    
    - **docs**: audits & other technical docs (not the reference) 
    
    - **ethClient**: the ethereum RPC client 
    
    - **ethdb**: the data store code goes there 
    
    - **ethstats**: network reporting service. 
    
    - **event**: Publish/Subscribe module for events 
    
    - **graphql**: graphql interface for node data
    
    - **internal**: Utility code to interact with the internal eth API
    
    - **les**: this package implements the Light Ethereum Subprotocol
    
    - **Light**: this package implements on-demand retrieval capable state and chain objects for the ethereum Light Client. 
    
    - **log**: Logging package 
    
    - **metric**: Forked module to create and update metrics
    
    - **miner**: This package implements Ethereum block creation and mining 
    
    
    - **mobile**: Android and IOS support (wrappers for an ethereum client)
    
    - **node**: sets up multi-protocol Ethereum nodes. In the model expressed by this package, a node is a collection of services which use shared resources to provide RPC APIs. Services can also offer devp2p protocols, which are wired up to the devp2p network when the node instance is started.

    - **P2P**: the package p2p implements the Ethereum p2p network protocols.  
    
    - **params**: network parameters that need to be constant between clients, but 
    aren't necessarilly consensus related. 

    - **rlp**: the package rlp implements the RLP (recursive linear prefix) serialization format. the rlp goal is to encode arbitrarily nested arrays of 
    binary data and is the main encoding method used to serialize objects in Ethereum.


    - **rpc**: this package implements bi-directional JSON-RPC 2.O on multiple
    transports, it provides access to the exported methods of an object accross a
    network or other I/O connection. After creating a server or client instances,
    object can be registered to make them visible as 'services'. Exported methods that
    follow specific conventions can be called remotely. 


    - **signer**: transaction validation, (this part should be explored in depth)



### Bitcoin Go



    - **addrmngr**: Concurrency safe bitcoin address manager 
    - **blockchain**: Blockchain handling and block selection rule 
    - **btcec**: eliptic curve cryptography needed for btc 
    - **btcjson**: Provides primitives for working with the Bitcoin Json-RPC api
    - **btcutils**: Provides bitcoin specific convenience functions and types 
    - **chaincfg**: Chain configuration parameters 
    - **cmd**: command line utilities 
    - **connmngr**: bitcoin network connection manager 
    -**database**: block and metadata storage 
    -**integration**: integration tests 
    -**mempool**: Provides a policy-enforced pool of unmined bitcoin transactions
    -**mining**: Handles bitcoin minging 
    -**netsync**: A concurrency safe block syncing protocol
    -**peer**: Provides a common base for creating and managing bitcoin network peers
    -**release**: contains the build scripts  that the btcd uses in order to build binaries for each new relases

    -**rpcclient**: this package implements a websocket enabled bitcoin JSON-RPC client

    -**txscript**: this package implements the bitcoin script language 

    -**wire**: this package implements the bitcoin wire protocol


## Features of our blockchain (Requirements)

    - Decentralized P2P network 
    - Store transactions in blocks 
    - the data store should be immutable 
    - the consensus algo should be robust (proof of work)
    - A native currency that is minted after every mining round 
    - the currency will have a max cap 
    - users should be able to deploy smart contracts on this blockchain 
    - Unanimous : all transactions should be transparent 
    - Should be able to connect to a wallet and send coins in it. 
    - users can interact with the blockchain data through an API 

## How blockains work 

### Ethereum 

    Ethereum is a permissionless non-hierarchical network of computers (nodes) that build and come to a consensus on a ever growing series of blocks , or batches  of transactions, known as the blockchain . each block contains an identifier  of the chain
    that must preced it if the block in the order they are listed, thereby alterring the
    ETH  balance 

    - Whenever a blocks adds to its chain, it executes the transactions in the block in the order they are listed, thereby altering the ETH balances and other storage values 
    of ethereum accounts. these balances and values collectively known as the state are 
    maintained separately from the blockchain on a merkle tree 


    - Each node communicates with a relatively small subset of the network -- it's peers. 
    whenever a node wishes to include a new transaction in the blockchain, it sends a copy 
    of the transaction to each of its peers, who then send a copy to each of their peers, and so on. 
    in this way it propagates through the network. Certain nodes called miners, maintain 
    a list of all these new transactions and use them to create new blocks, which they then send to the rest of the network. whenever a node  receives a block, it checks 
    the validity of the block and all of the transactions therein and if it finds the block
    to be valid, adds it to its blockchain and executes all of these transactions.

    Since block creation and broadcasting are permissionless, a node may receive multiple 
    blocks competing to be successor to a particular block. the node keep tracks of all the
    valid chains that result from it and regularly drops the shortest one. 

    According to the ethereum protocol, the longest chain at any given time is to be considered the canonical one.


    ## Ether 

    - Ether or ETH is the cryptocurrency generated in accordance with the ethereum protocol
    as a reward to miners in a proof-of-work system for adding blocks to the blockchain. 
    Ether is represented in the state as an unsingned integer associated with each account
    this being the account's,  ETH balance denominated in Wei (10**18 wei = 1 ETH). each 
    block, new ETH is generated by the addition of a protocol-specified amount, currenlty 
    2x10**18 Wei (2 ETH), to the balance of the account of the miner's choosing. this is 
    known as the block reward.  

    - Additonaly, ether is the only currency accepted by the protocol as payment for  
    a transaction fee, which also goes to the miner. the block reward together with  
    the transaction fee

    - The block reward together with the transaction fee provide incentives to the 
    miners to keep the blockchain growning (i.e to keep processing new transactions). 
    therefore , ETH is fundamental to the operation of the network.

    - Ether may be sent from one account to another wich may only entails subtracting from 
    one account balance and adding  another accounts balance

    - The shift to ether 2.0 may reduce the issuance rate of ether. there is currently no 
    implemented hard cap on the total supply of ether

    ## Accounts 

    There are two types of accounts on Ethereum: user accounts (also known as
    externally-owned-accounts) and contracts. both types have an ETH balance, may 
    send ETH to any account, may call any public function of a contract to create a
    new contract, and are identified on the blockchain and in the state by an account
    address. 

    User accounts are the only type of account that may create transactions. For a transaction to be valid it must be signed using the sending account's private key
    the 64 character hexadecimal string from which the account's address is derived 
    
    - The algorithm used to produce the signature is ECDSA (Elliptic curve digital
    signature algorithm). importantly this algorithm allows one to derives the signer address from the signature 
    withouth knowing the private key.


    - Contracts are the only type of account that has associated code ( a set of function and variable declarations) 
    and a contract storage ( the values of variables at any given time). A contract function may take arguments  
    and may have return values. in addition to control flow statements, the body of a function may include 
    instructions to ETH, read from and write the contract's storage, create temporary storage (memory) that 
    vanishes at the end of the function, perform arithmetic and hashing operations, call contract's own functions
    call public functions of other contracts, create new contracts, and query information about the current transaction
    or the blockchain. 


    ## Addresses 

    - Ethereum addresses are composed of a prefix "0x" (a common identifier for hexadecimal) concatenated 
    with the rightmost 20 bytes of the Keccak-256 hash of the ECDSA public key (the curve used is so the 
    so called secp256k1 ). In hexadecimal two digits represent a byte, and so addresses contain 40 hexadecimal
    digits (e.g 0xb794f5ea0ba39494ce839613fffba74279579268 )


    ## Virtual Machine 

    - The ethereum virtual machine is the runtime environment for transaction execution 
    in Ethereum. it includes  a stack, memory, gas balance, program counter and the 
    persistent storage for all accounts including contract code. 

    - When a transaction calls a contract function, the arguments in the call are 
    added to the stack and the EVM translates the contract's function, the argument in 
    the call are added to the stack and the EVM translates the contract's bytecode into 
    stack operations. 

    - Stack items may stored in memory or storage, and data from memory/storage may 
    be added to the stack.

    - The EVM is isolated from the other files and processes on the node's compute to ensure 
    that for a given pre-transaction state and transaction, every node produces the same 
    post transaction state, thereby enabling network consensus.



    ## Gas 


    Gas is a unit of account within the EVM used in the calculation of a transactio fee. which is 
    the amount of ETH a transaction's sender must pay to the miner who includes the transaction in 
    the blockchain. 

    Each type of operation which pay may be performed by the EVM is hardcoded with a certiain gas cost, 
    which is intended to be roughley proportional to the amount of resources. computation and storage
    a node must expend to perform that operation. 

    when a sender creates a transaction, the sender must specify a gas limit and gas price. the gas 
    limit is the maximum amount of gas the sender is willing to use in the transaction, and the gas 
    price, the more incentive a miner has to include the transaction in their block, and thus the quicker
    the transaction will be included in the blockchain. 

    The sender buys the full amount of gas upfront (i.e their ETH balance is debited the amount: 
    gas limit * gas_price), at the start of the execution of the transaction, and is refunded at the 
    end for any unused gas. if at any point the transaction does not have enough gas to perform the 
    next operation, the transaction is reverted but the sender is only refunded for the unused gas. 

    Governance 

    In 2015 a development governance was proposed as the ethereum improvement propsal (EIP), standardized
    on EIP-1. the core development group and community were to gain consensus by a process-regulated EIP 


    ## Difficulty Bomb 

    The difficulty bomb is an Ethereum protocol feature that causes the difficulty of mining a block  
    to increase exponentially over time after a certain block is reached, with the intended purpose 
    being to incentivize upgrades to the protocol and prevent miners from having too much control over 
    upgrades. As the protocol is upgraded the difficulty bomb is typically pushed further out in time. 

    It was originally placed there primarily to ensure a successful upgrade from proof of work to proof of 
    stake, an upgrade that removes miners entirely from the design of the network. the period during 
    which the mining difficulty is increased is called the "Ice Age" 

    ETH VS BITCOIN 

    BTC primary use case is a store of value and a digital currency. Ether can also be used as a digital 
    currency and store of value, but the ethereum network also makes it possible to create and run 
    decentralized applications and smart contracts. 

    Blocks in ethereum are validated every 12 seconds while it takes ~10 minutes to validate a block on the 
    BTC network.  


    ## Applications 

    The EVM instruction set is Turin-Complete. popular uses of ethereum have included the creation of 
    fungible (ERC20) and non-fungible (ERC721) tokens with a variety of properties. 

    ## Contract Source code 

    Ethereum's smart contracts are written in high-level programming languages and then compiled down to 
    EVM bytecode and deployed to the ethereum blockchain. (Solidity, Serpent, Yul, Vyper) 

    Source code and compiler information are usually published along with the launch of the contract so 
    that users can see the code and verify that it compiles to the bytecode that is on-chain. 

    One issue related to using smart contracts on a public blockchain is that bugs, including security 
    holes, are visible to all but cannot be fixed quickly.


### Bitcoin network 

    To be taken later 



### Core Features to be implemented 

    - The blockchain node 
        - native currency
        - transactions 
        - consensus (proof of work) 
        - smart contracts?
        - P2P mechanisms 


        


