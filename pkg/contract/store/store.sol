// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract GymMembership is Initializable {

    address public owner;
    uint256 public membershipCounter;

    enum MembershipStatus {
        None,
        Active,
        Inactive
    }

    struct Membership {
        uint256 id;
        string email;
        string membershipType;
        uint256 startDate;
        uint256 endDate;
        MembershipStatus status;
    }

    struct User {
        string name;
        string email;
        string contact;
        uint256 membershipId;
    }

    mapping(string => User) public users;
    mapping(uint256 => Membership) public memberships;

    event UserRegistered(
        string email,
        string name,
        string contact
    );
    event UserVerified(string email);
    event MembershipPurchased(string email, uint256 membershipId, string membershipType, uint256 startDate, uint256 endDate);
    event MembershipActivated(string email, uint256 membershipId);
    event MembershipCancelled(string email, uint256 membershipId);
    event MembershipRenewed(string email, uint256 membershipId, uint256 newEndDate);
    event UserDeleted(string email);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action.");
        _;
    }

    modifier onlyRegisteredUser(string memory _email) {
        require(
            bytes(users[_email].email).length != 0,
            "Only registered user can perform this function."
        );
        _;
    }

    function initialize() public initializer {
        owner = msg.sender;
        membershipCounter = 0;
    }

    function registerUser(
        string memory _name,
        string memory _email,
        string memory _contact
    ) public onlyOwner{
        require(
            bytes(users[_email].email).length == 0,
            "User already registered"
        );

        users[_email] = User({
            name: _name,
            email: _email,
            contact: _contact,
            membershipId: 0
        });

        emit UserRegistered(_email, _name, _contact);
    }

    function verifyUser(string memory _email) public onlyOwner returns (bool){
        require(
            bytes(users[_email].email).length != 0,
            "User is not registered"
        );

        emit UserVerified(_email);
        return true;
    }

    function purchaseMembership(
        string memory _email,
        string memory _membershipType,
        uint256 _startDate,
        uint256 _endDate
    ) public onlyRegisteredUser(_email) {
        uint256 userMembershipId = users[_email].membershipId;
        Membership storage userMembershipDetails = memberships[userMembershipId];
        require(
            userMembershipId == 0 || userMembershipDetails.status != MembershipStatus.Active,
            "Active membership exists, cancel or renew it before purchasing a new one"
        );
        require(_startDate >= block.timestamp, "Start date must be in the future");
        require(_endDate > _startDate, "End date must be after start date");

        membershipCounter++;
        memberships[membershipCounter] = Membership({
            id: membershipCounter,
            email: _email,
            membershipType: _membershipType,
            startDate: _startDate,
            endDate: _endDate,
            status: MembershipStatus.Inactive
        });

        users[_email].membershipId = membershipCounter;

        emit MembershipPurchased(_email, membershipCounter, _membershipType, _startDate, _endDate);
    }

    function activateMembership(uint256 _membershipId) public onlyOwner {
        Membership storage membership = memberships[_membershipId];
        require(bytes(membership.email).length != 0, "Membership does not exist");
        require(
            membership.status == MembershipStatus.Inactive,
            "Membership is already active"
        );

        membership.status = MembershipStatus.Active;
        emit MembershipActivated(membership.email, _membershipId);
    }

    function cancelMembership(uint256 _membershipId) public onlyOwner {
        Membership storage membership = memberships[_membershipId];
        require(bytes(membership.email).length != 0, "Membership does not exist");
        require(
            membership.status == MembershipStatus.Active,
            "Membership is not active"
        );

        membership.status = MembershipStatus.Inactive;

        emit MembershipCancelled(membership.email, _membershipId);
    }

    function renewMembership(uint256 _membershipId, uint256 _newEndDate)
        public
        onlyOwner
    {
        Membership storage membership = memberships[_membershipId];
        require(bytes(membership.email).length != 0, "Membership does not exist");
        require(membership.status == MembershipStatus.Active, "Membership is not active");
        require(_newEndDate > membership.endDate, "New end date must be after current end date");
        membership.endDate = _newEndDate;

        emit MembershipRenewed(membership.email, _membershipId, _newEndDate);
    }

    function checkMembershipStatus(string memory _email)
        public
        view
        returns (MembershipStatus)
    {
        uint256 membershipId = users[_email].membershipId;
        if (membershipId == 0) {
            return MembershipStatus.None;
        }
        return memberships[membershipId].status;
    }

    function getMembershipDetails(uint256 _membershipId)
        public
        view
        returns (
            string memory,
            string memory,
            uint256,
            uint256,
            MembershipStatus
        )
    {
        Membership storage membership = memberships[_membershipId];
        return (
            membership.email,
            membership.membershipType,
            membership.startDate,
            membership.endDate,
            membership.status
        );
    }

     function deleteUser(string memory _email) public onlyOwner {
        require(bytes(users[_email].email).length != 0, "User does not exist");
        delete memberships[users[_email].membershipId];
        delete users[_email];
        emit UserDeleted(_email);
    }
}
