class Pagination {
  currentPage = 0;

  perPage = 0;

  lastPage = 0;

  from = 0;

  to = 0;

  total = 0;

  constructor(
    currentPage: number = 1,
    perPage: number = 25,
    lastPage: number = 1,
    from: number = 0,
    to: number = 0,
    total: number = 0,
  ) {
    this.currentPage = currentPage;
    this.perPage = perPage;
    this.lastPage = lastPage;
    this.from = from;
    this.to = to;
    this.total = total;
  }

  hasPage(page: number): boolean {
    return (page >= 1 && page <= this.lastPage);
  }

  isOnFirst(): boolean {
    return this.currentPage === 1;
  }

  isOnLast(): boolean {
    return this.currentPage === this.lastPage;
  }

  count(): number {
    return this.getFrom() - this.getTo();
  }

  getFrom(): number {
    return this.from;
  }

  getTo(): number {
    return this.to;
  }

  getTotal(): number {
    return this.total;
  }

  isEmpty(): boolean {
    return this.total <= 0;
  }

  getPageFromDirection(direction = 'next'): number {
    let page = this.currentPage;
    if (direction === 'next') {
      if (this.lastPage > page) {
        page += 1;
      }
    } else if (direction === 'prev') {
      if (page > 1) {
        page -= 1;
      }
    }

    return page;
  }
}

export default Pagination;
