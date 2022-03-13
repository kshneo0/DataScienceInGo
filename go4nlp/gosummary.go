package main

import (
	"fmt"

	"github.com/JesusIslam/tldr"
)

func main() {
	// mydocx := `The history of Artificial Intelligence (AI) began in antiquity, with myths, stories and rumors of artificial beings endowed with intelligence or consciousness by master craftsmen. The seeds of modern AI were planted by classical philosophers who attempted to describe the process of human thinking as the mechanical manipulation of symbols. This work culminated in the invention of the programmable digital computer in the 1940s, a machine based on the abstract essence of mathematical reasoning. This device and the ideas behind it inspired a handful of scientists to begin seriously discussing the possibility of building an electronic brain.

	// The field of AI research was founded at a workshop held on the campus of Dartmouth College during the summer of 1956.[1] Those who attended would become the leaders of AI research for decades. Many of them predicted that a machine as intelligent as a human being would exist in no more than a generation, and they were given millions of dollars to make this vision come true.

	// Eventually, it became obvious that they had grossly underestimated the difficulty of the project. In 1973, in response to the criticism from James Lighthill and ongoing pressure from congress, the U.S. and British Governments stopped funding undirected research into artificial intelligence, and the difficult years that followed would later be known as an "AI winter". Seven years later, a visionary initiative by the Japanese Government inspired governments and industry to provide AI with billions of dollars, but by the late 80s the investors became disillusioned and withdrew funding again.

	// Investment and interest in AI boomed in the first decades of the 21st century when machine learning was successfully applied to many problems in academia and industry due to new methods, the application of powerful computer hardware, and the collection of immense data sets.`

	mydocx := `인공지능의 역사는 20세기 초반에서 더 거슬러 올라가보면 이미 17~18세기부터 태동하고 있었지만 이때는 인공지능 그 자체보다는 뇌와 마음의 관계에 관한 철학적인 논쟁 수준에 머물렀다. 그럴 수밖에 없는 것이 당시에는 인간의 뇌 말고는 정보처리기계가 존재하지 않았기 때문이다. 그러나 시간이 흘러 20세기 중반부터 본격적으로 컴퓨터 발달 혁신의 물줄기가 터지기 시작하면서 이거 잘하면 컴퓨터로 두뇌를 만들어서 우리가 하는 일을 시킬 수 있지 않을까? 라는 의견이 제시되었고 많은 사람들이 그럴 듯하게 여겨 빠른 속도로 인공지능은 학문의 영역으로 들어서기 시작했다.

인공지능(AI, Artificial Intelligence)이라는 용어가 처음 등장한 때는 1956년에 미국 다트머스 대학교에서 마빈 민스키, 클로드 섀넌 등 인공지능 및 정보 처리 이론에 지대한 공헌을 한 사람들이 개최한 학회에서 존 매카시가 이 용어를 사용하면서부터이다. 하지만 인공지능이라는 개념 자체는 훨씬 예전부터 있었다. 예를 들면, 앨런 튜링이 ‘생각하는 기계’의 구현 가능성과 튜링 테스트를 제안한 것은 1950년의 일이며, 최초의 신경망 모델은 1943년에 제안되었다.[1]

그리고 당연하지만 이런 일에 관심을 가진건 서방뿐만이 아니어서, 소련 역시 아나톨리 키토프 박사가 본인의 저서 "붉은 서"에서 "ЕГСВЦ(Единой централизованной автоматизированной системы управления народным хозяйством страны - 국가(계획)경제 네트워크 중심적 통제체계)" 라는 것을 제시하였는데, 이는 컴퓨터 네트워크화를 통한 더 나은 계획 경제 체제와 사회의 추구를 목표로 삼은 이론이었다. 이것을 소련 컴퓨터 공학자 빅토르 글루쉬코프가 더욱 개량한 것이 바로 OGAS(ОГАС - Общегосударственная автоматизированная система учёта и обработки информации, 전연방자동정보처리체계) 계획이다.

20세기 중반에도 인공지능 연구는 자연어처리나 복잡한 수학 문제를 해결하는 등 정말로 인간만이 할 수 있는 영역에 있던 문제들도 컴퓨터로 착착 해결할 수 있었던 상당히 혁신적인 연구였으며, 많은 관심을 받고 지속적으로 연구가 이루어진 분야이다. 당연히 AI 산업은 이미 1980년도에 10억불 규모의 시장을 형성할 정도로 큰 분야였으므로 과거에 이런저런 이유로 관심이 없었다던가 실용화가 되지 않았다는 것은 어불성설이다. 다만 아무래도 당시의 정보처리 능력의 한계와 정보량의 부족, 그리고 이런저런 이유로 연구자금지원이 중단되는 트러블과 특히 1969년도에 Marvin Minsky와 Seymour Papert가 "Perceptrons"이라는 책을 출간하면서 지적한 단일 계층 신경망의 한계로 인해 1970년대에 한동한 인기가 시들시들 하기도 했었다.(1차 AI 겨울/AI winter) 이 문제는 1980년도에 다층 신경회로망이 도입되면서 해소되었지만 정보처리 능력의 한계와 해소되기까지는 더 시간이 필요했다.

이후 1974년도에 제시된 역전파 알고리즘, 전문가시스템의 성장과 1980년도에 신경망 이론에 대한 연구가 다시 재개되면서 많은 연구가 있었지만 여전히 성장이 지지부진하여 큰 실망을 안겨주기도 했다.(2차 AI 겨울/AI winter) 문자인식이나 음성인식등의 가시적인 성과가 있는 분야도 있었지만 대화 인공지능등의 개발 실패 등, 눈앞의 목표를 달성하지 못하는 경우도 많았기 때문인데, 심지어 이런 부분은 수십년이 지난 현재도 극복하지 못해서 아직까지는 인간과 대화를 한다기보다는 자동 응답기에 가까운 수준이다. 이 때문에 1990년도 이후부터 인공지능의 목표는 인간지능의 구현이라는 막연히 넓은 목표에서 문제해결과 비즈니스 중심으로 더 신중하고 좁은 분야가 되었으며, 그제서야 때맞춰 나타난 하드웨어의 성장을 업고 더 성공적인 분야가 될 수 있었다.`

	summary := tldr.New()
	results, _ := summary.Summarize(mydocx, 1)

	fmt.Println(results)

}
