
dirs = 	testppack \
	arrays \
	chanloop \
	chanrange \
	chanasyn \
	chansvr1 \
	chansvr2 \
	chansyn \
	chantest \
  defer \
  envshow \
  errtest \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
	flagdemo \
  force_type \
	func_lit \
	func_ret \
	gorsynchs/chanex \
	gorsynchs/mutexex \
	gorsynchs/wgroupex \
	gortns01 \
	gt044 \
	hello \
	hello_small \
	httpserv \
	interface_01 \
	interface_02 \
	interface_03 \
	list \
	list_struct \
	loghello \
	maps \
	maps_ss \
	maps_merge \
	methods \
	multisrc \
	netcon \
	numconv \
	oltesta \
	ranges \
	rand_between \
	shax \
	showgo \
	showmac \
	showtime \
	sliceappt \
	slices \
	stringlens \
	strings_utils \
	struct \
	switch_demo \
	symlshow \
	symltest \
	tcp01 \
	tcp02 \
	tcp03 \
	tcp04 \
	tcp05 \
	tcp06 \
	type_String \
	unicloup \
	uniques \
	utfconv \


packlist =	numbers \

.PHONY: $(dirs) $(packlist) packages clean format

all: $(dirs)
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go build; \
	cd $$curd; \
	done


packages:
	@for i in $(packlist); do \
	echo $$i; \
	go install -v go-samp/$$i; \
	done

clean:
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go clean; \
	cd $$curd; \
	done

format:
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && gofmt -w -tabwidth=2 *.go; \
	cd $$curd; \
	done

