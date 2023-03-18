package p1472

type BrowserHistory struct {
	h     []string
	index int
	end   int
}

func Constructor(homepage string) BrowserHistory {
	urls := make([]string, 5005)
	urls[0] = homepage
	return BrowserHistory{
		h:     urls,
		index: 0,
		end:   0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.index++
	this.h[this.index] = url
	this.end = this.index
}

func (this *BrowserHistory) Back(steps int) string {
	this.index -= steps
	if this.index < 0 {
		this.index = 0
	}
	return this.h[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.index += steps
	if this.index > this.end {
		this.index = this.end
	}
	return this.h[this.index]
}
