package hwclock

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

// HwclockTimeToTimeEl7 hwclock 时间转化成标准时间
func HwclockTimeToTimeEl7(args string) (time.Time, error) {
	var (
		reply time.Time
		err   error
	)

	//麒麟环境
	//[root@node2 ~]# hwclock -r
	//2023-03-21 10:38:59.995906+08:00
	//2023-07-31 23:59:10.000904+08:00
	//2023-08-01 00:00:11.018690+08:00

	//centos环境
	//[root@node1 ~]# hwclock -r
	//Tue 21 Mar 2023 02:40:23 AM CST  -0.650751 seconds
	//Mon 31 Jul 2023 12:59:53 AM EDT  -0.985555 seconds
	//Mon 31 Jul 2023 01:00:54 AM EDT  -0.985621 seconds
	//Mon 31 Jul 2023 12:59:23 PM EDT  -0.985565 seconds
	//Mon 31 Jul 2023 01:00:24 PM EDT  -1.001135 seconds

	stdoutList := strings.Split(args, " ")
	if len(stdoutList) < 7 {
		logrus.Errorf("len(stdoutList) < 7. stdoutList:%+v", stdoutList)
		return reply, fmt.Errorf("len(stdoutList) < 7. stdoutList:%+v", stdoutList)
	}

	//  Tue 21 Feb 2023 09:01:55 AM IST  -0.559689 seconds
	weekDayStr := stdoutList[0]
	dayStr := stdoutList[1]
	monthStr := stdoutList[2]
	yearStr := stdoutList[3]
	timeStr := stdoutList[4]
	pmStr := stdoutList[5]
	timezoneStr := stdoutList[6]
	if len(strings.Split(timeStr, ":")) < 3 {
		logrus.Errorf("len(timeStr) < 3. stdoutList:%+v", stdoutList)
		return reply, fmt.Errorf("len(timeStr) < 3. stdoutList:%+v", stdoutList)
	}
	hourStr := strings.Split(timeStr, ":")[0]
	minuteStr := strings.Split(timeStr, ":")[1]
	secondStr := strings.Split(timeStr, ":")[2]
	if pmStr == "PM" {
		hourInt, err := strconv.Atoi(hourStr)
		if err != nil {
			logrus.Errorf("strconv.Atoi err:%v, hourStr:%s", err, hourStr)
			return time.Time{}, err
		}
		if hourInt >= 12 {
			hourInt = hourInt - 12
		}

		hourInt += 12
		hourStr = strconv.Itoa(hourInt)
		timeStr = fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
	} else {
		hourInt, err := strconv.Atoi(hourStr)
		if err != nil {
			logrus.Errorf("strconv.Atoi err:%v, hourStr:%s", err, hourStr)
			return time.Time{}, err
		}
		if hourInt >= 12 {
			hourInt = hourInt - 12
		}

		hourStr = strconv.Itoa(hourInt)
		timeStr = fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
	}

	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}

	// Tue Feb 02 09:01:55 IST 2023
	dateStr := fmt.Sprintf("%s %s %s %s %s %s", weekDayStr, monthStr, dayStr, timeStr, timezoneStr, yearStr)
	reply, err = time.Parse("Mon Jan _2 15:04:05 MST 2006", dateStr)
	if err != nil {
		logrus.Errorf("time.Parse: err:%v, dateStr:%s", err, dateStr)
		return reply, err
	}
	return reply, err
}

