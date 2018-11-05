![Schema](/schema.svg)


# student
a student is a person that enrolls in a class.

__Properties__
- ID _uuid pk_
- Firstname _string_
- Lastname _string_
- Birthdate _time_
- Email _string_
- StateID _string_


# Teacher
A tacher is a person who instucts student in learning. They are highly quallified indiuviuals. Schools choose Teachers that can fullfill thier Curriculum objectives.

__Properties__
- ID _uuid pk_
- Firstname _string_
- Lastname _string_
- Birthdate _time_
- Email _string_
- StateID _string_

# Curriculum
A Curriculum defines a field of study, disipline, or trade in the most abstract sense. A Curriculum should highlight the objectives or results it will produce in a student. A Curriculum does not include the exact manner is which it is taught. The implementation of a Curriculum shoud be defined by a Syllabus. A Pincipal or Office Administrator will often dictate the Curriculums a school will offer.

__Properties__
- ID _uuid pk_
- Name _string_
- Description _string_


# Syllabus
A Syllabus describes the way a Curriculum is taught. A Syllabus should define the work a student should complete in order to fullFill the objectives the Curriculum. A Teacher or group of Teachers are often deeply involved in defining a Syllabuses implementation.

__Properties__
- ID _uuid pk_
- Curriculum _fk_
- Name
- Description


# District
A district is business entity that manages school. They often are responsible for staffing and creating guildlines for thier school.

__Properties__
- ID _uuid pk_
- Name


# School
A School is an entity for learning in the most abstract sense. Schools often have specific objectives that they would like there students to acheive. These objectives are often accompanied by specific styles of teaching. The Courses a School offers usually is a reflection of its objectives.

__Properties__
- ID _uuid pk_
- District _fk_
- Name


# Campus
A campus is a physical location and is addressable. A Campus usually belongs to a single school. a single school can have many campuses. A Campus can also be online, like a specific web domain. 

__Properties__
- ID _uuid pk_
- School _fk_
- Name
- Address


# Room
A Room belongs to a Campus. A Room is a single place a student would expect to go to attend a class, receive intruction from a teacher, and do class work. A Room can also be online, like a classroom URL. 

__Properties__
- ID _uuid pk_
- Campus _fk_
- Name


# Period
A Period is a duration of time durring in a day. Perids reprents time slots that a class can be held. Periods allow classes to be scheduled.

__Properties__
- ID _uuid pk_
- Campus _fk_
- Name
- StartTime
- EndTime


# Calendar
A Calendar primarily it is used to list school days. but it can be used to track any school event.

__Properties__
- ID _uuid pk_
- Campus _fk_
- Event _look_
- Date


# Class
A Class is an instance of a syllubas that is eventually scheduled and filled with students. A Student enrolls in a Curriculum by being assigned to a Class. Classes can be schedule before students enroll if needed. A class is intended to organize students into group sizes small enough to fit into a class room or meet teacher and student ratios. when scheduling a class, the time it takes for a student to complete the syllabus and fullfill the expectations of the Curriculum should be taken into consideration. Class should be schdeuled 1 or more times for students to attend it.

__Properties__
- ID _uuid pk_
- Syllabus _fk_
- Curriculum _fk_
- MaxSize _int_
- Teacher _fk_


# Schedule
A Schedule is a list of Classes and the time and place they will be held. 

__Properties__
- ID _uuid pk_
- Date _fk_
- Period _fk_
- Room _fk_
- Class _fk_


# Session
Sessions are historical instances of Classes being "in session". Sessions are created at the time a class is "in session" and they are event driven. a sessions should link back to a Schedule event. A session can not be deleted once attendance has been taken fot it.

__Properties__
- ID _uuid pk_
- Schedule _fk_
- Teacher


# Assignment
An Assignment is any task to peice of work that a student would be expected to complete to receive a grade in a class. An Assignment is a list of questions and Answer types, or it can be a Task this is just given a score.

__Properties__
- ID _uuid pk_
- Session _fk_
- Name 



# Question
Assignments can have 1 or more questions. the type of question will detirm the type of Answer.

__Properties__
- ID _uuid pk_
- Assignment _fk_
- Text
- Type


# Answer
An Answer corresponds to a students answer to an Assignment Question.

__Properties__
- ID _uuid pk_
- student _fk_
- Question _fk_
- Timestamp _time_
- Response _?_


# Attendnace
Attendance should be captured for every enrolled student in every session. attenance is stored as a percentage of the time a student spent in class. 1 being the entire class and 0 being absent. a time percentage will allow for clerical errors in the schedule to be corrected as long as the corrections reflect real life events.

__Properties__
- ID _uuid pk_
- Session _fk_
- Student _fk_
- Percentage _float_
