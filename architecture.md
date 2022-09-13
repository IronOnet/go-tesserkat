### System design 

        # Steps: 
            1. Understand the problem and establish a design scope 
            2. Propose a high level design 
            3. Design deep dive 


        - Actors: 
            - Users of the cryptocurrency 
            - contract developers 
            - miners 


       

    

## Analysis 

    - Actors 
        - normal users with wallets and their addresses 
        - Miners 


    ## Components 
        - Nodes 
        - Wallets 


        => Nodes: 
            - store the transaction data 
            - validate chains through proof of work 
            - add blocks to the blockchain 
            - send copies of the new state to other nodes 
            - compete with other blocks for block rewards 
            - stores the smart contracts 


        => Wallet: 
            - contains addresses 
            - private and public keys of the wallet 
            - contain the balance of the wallet 



    ## Components Low level 


        ## Node 

            - cmd : Command-line utilities 
            - p2p: peer management 
            - database: database managment 
            - mempool: mempool 
            - blockchain: Blockchain handling selection rule  
            - core: virtual machine 
            - tests: Tests 
            - sync: synchronization 
            - mining: block creation and mining 
            - graphql: graphql interface for node data 
            - json: json interface for node data
            - log: logs handling package 
            - consensus: consensus mechanisms 
            - 

