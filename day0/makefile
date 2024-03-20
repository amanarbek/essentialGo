TEST1 = --Первый тест--
TEST2 = --Второй тест--
TEST3 = --Третий тест--
TEST4 = --Четвертый тест--
TEST5 = --Пятый тест--
PATHS  = tests/

make:
	@go build EX_00.go

tests:
	@$(MAKE) test1
	@echo
	@$(MAKE) test2
	@echo
	@$(MAKE) test3
	@echo
	@$(MAKE) test4
	@echo
	@$(MAKE) test5
	@echo "Всё чики пуки ёпта!"
	
test1:
	@echo $(TEST1)
	@cat $(PATHS)test1.txt | ./EX_00

test2:
	@echo $(TEST2)
	@cat $(PATHS)test2.txt | ./EX_00

test3:
	@echo $(TEST3)
	@cat $(PATHS)uniqTest3.txt | ./EX_00

test4:
	@echo $(TEST4)
	@cat $(PATHS)uniqTest4.txt | ./EX_00

test5:
	@echo $(TEST5)
	@cat $(PATHS)uniqTest5.txt | ./EX_00

clean:
	@rm -f EX_00

.PHONY: tests