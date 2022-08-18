###
# Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],…] (si < ei),
# find the minimum number of conference rooms required.
#
# For example, Given [[0, 30],[5, 10],[15, 20]], return 2.
###
import heapq


def find_meeting_rooms(meetings):
    rooms = []
    heapq.heappush(rooms, 0)

    for x, y in meetings:
        available = False
        if rooms[0] <= x:
            available = True
            heapq.heappop(rooms)
            heapq.heappush(rooms, y)

        if not available:
            heapq.heappush(rooms, y)

    return len(rooms)


if __name__ == '__main__':
    print(find_meeting_rooms([(0, 30), (5, 10), (15, 20)]))
    print(find_meeting_rooms([(0, 30), (5, 10), (7, 8), (15, 20)]))
