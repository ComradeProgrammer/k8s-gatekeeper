import os
from framework.TestSuite import TestSuite
def main():
    testPath = os.path.dirname(__file__)
    workspacePath = os.path.abspath(testPath+"/..")
    

    dirs=os.listdir(workspacePath+"/example")
    for i in range(0,len(dirs)):
        testsuit=TestSuite(dirs[i])
        res=testsuit.run()
        print(res)

if __name__ == '__main__':
    main()