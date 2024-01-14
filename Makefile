.PHONY: adr dd commit 


##### scaffing

ADR_COUNT:=$(shell find docs/ADR -type f | wc -l | tr -d ' ') 
DD_COUNT:=$(shell find docs/DesignDog -type f | wc -l | tr -d ' ') 
adr:
	npx scaffdog generate ADR --output 'docs/ADR' --answer 'number:${ADR_COUNT}'

dd:
	npx scaffdog generate DD --output 'docs/DesignDog' --answer 'number:${DD_COUNT}'

##### git

commit:
	npx git-cz



